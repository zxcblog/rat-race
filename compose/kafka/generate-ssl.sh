#!/usr/bin/env bash

set -e

KEYSTORE_FILENAME="kafka.keystore.jks"
VALIDITY_IN_DAYS=3650
DEFAULT_TRUSTSTORE_FILENAME="kafka.truststore.jks"
TRUSTSTORE_WORKING_DIRECTORY="truststore"
KEYSTORE_WORKING_DIRECTORY="keystore"
CA_CERT_FILE="ca-cert"
KEYSTORE_SIGN_REQUEST="cert-file"
KEYSTORE_SIGN_REQUEST_SRL="ca-cert.srl"
KEYSTORE_SIGNED_CERT="cert-signed"
SUBJ="/CN=localhost/OU=TEST/O=CONFLUENT/L=PaloAlto/ST=Ca/C=US"
PASSWORD="confluent"


function file_exists_and_exit() {
  echo "'$1' 不能存在。在之前移动或删除它"
  echo "重新运行此脚本。"
  exit 1
}

if [ -e "$KEYSTORE_WORKING_DIRECTORY" ]; then
  file_exists_and_exit $KEYSTORE_WORKING_DIRECTORY
fi

if [ -e "$CA_CERT_FILE" ]; then
  file_exists_and_exit $CA_CERT_FILE
fi

if [ -e "$KEYSTORE_SIGN_REQUEST" ]; then
  file_exists_and_exit $KEYSTORE_SIGN_REQUEST
fi

if [ -e "$KEYSTORE_SIGN_REQUEST_SRL" ]; then
  file_exists_and_exit $KEYSTORE_SIGN_REQUEST_SRL
fi

if [ -e "$KEYSTORE_SIGNED_CERT" ]; then
  file_exists_and_exit $KEYSTORE_SIGNED_CERT
fi

echo
echo "欢迎使用Kafka SSL密钥库和信任库生成器脚本。"

echo
echo "首先，你需要生成一个信任存储和相关的私钥吗，"
echo "或者您已经有了信任存储文件和私钥？"
echo
echo -n "是否需要生成信任存储和关联的私钥？[y/n] "
read generate_trust_store

trust_store_file=""
trust_store_private_key_file=""

if [ "$generate_trust_store" == "y" ]; then
  if [ -e "$TRUSTSTORE_WORKING_DIRECTORY" ]; then
    file_exists_and_exit $TRUSTSTORE_WORKING_DIRECTORY
  fi

  mkdir $TRUSTSTORE_WORKING_DIRECTORY
  echo
  echo "好的，我们将生成一个信任存储和相关的私钥。"
  echo
  echo "首先，私钥。"
  echo
  echo "系统将提示您："
  echo " - 私钥的密码。记住这一点。"
  echo " - 有关您和您的公司的信息。"
  echo " - 注意，通用名称（CN）目前并不重要。"

#

  echo "正在执行生成证书命令"
  echo "openssl req -new -x509 -keyout $TRUSTSTORE_WORKING_DIRECTORY/ca-key \\"
  echo "-out $TRUSTSTORE_WORKING_DIRECTORY/$CA_CERT_FILE -days $VALIDITY_IN_DAYS \\"
  echo "-subj $SUBJ -passin pass:$PASSWORD -passout pass:$PASSWORD"

  openssl req -new -x509 -keyout $TRUSTSTORE_WORKING_DIRECTORY/ca-key \
    -out $TRUSTSTORE_WORKING_DIRECTORY/$CA_CERT_FILE -days $VALIDITY_IN_DAYS \
    -subj $SUBJ -passin pass:$PASSWORD -passout pass:$PASSWORD

  trust_store_private_key_file="$TRUSTSTORE_WORKING_DIRECTORY/ca-key"

  echo
  echo "Two files were created:"
  echo " - $TRUSTSTORE_WORKING_DIRECTORY/ca-key -- the private key used later to"
  echo "   sign certificates"
  echo " - $TRUSTSTORE_WORKING_DIRECTORY/$CA_CERT_FILE -- the certificate that will be"
  echo "   stored in the trust store in a moment and serve as the certificate"
  echo "   authority (CA). Once this certificate has been stored in the trust"
  echo "   store, it will be deleted. It can be retrieved from the trust store via:"
  echo "   $ keytool -keystore <trust-store-file> -export -alias CARoot -rfc"

  echo
  echo "Now the trust store will be generated from the certificate."
  echo
  echo "You will be prompted for:"
  echo " - the trust store's password (labeled 'keystore'). Remember this"
  echo " - a confirmation that you want to import the certificate"


  docker run --rm -it -v "$(pwd):/workspace" --workdir /workspace openjdk:latest keytool \
    -noprompt -keystore $TRUSTSTORE_WORKING_DIRECTORY/$DEFAULT_TRUSTSTORE_FILENAME \
    -alias CARoot -import -file $TRUSTSTORE_WORKING_DIRECTORY/$CA_CERT_FILE \
    -storepass $PASSWORD -keypass $PASSWORD

  trust_store_file="$TRUSTSTORE_WORKING_DIRECTORY/$DEFAULT_TRUSTSTORE_FILENAME"

  echo
  echo "$TRUSTSTORE_WORKING_DIRECTORY/$DEFAULT_TRUSTSTORE_FILENAME was created."

  # don't need the cert because it's in the trust store.
  rm $TRUSTSTORE_WORKING_DIRECTORY/$CA_CERT_FILE
else
  echo
  echo -n "Enter the path of the trust store file. "
  read -e trust_store_file

  if ! [ -f $trust_store_file ]; then
    echo "$trust_store_file isn't a file. Exiting."
    exit 1
  fi

  echo -n "Enter the path of the trust store's private key. "
  read -e trust_store_private_key_file

  if ! [ -f $trust_store_private_key_file ]; then
    echo "$trust_store_private_key_file isn't a file. Exiting."
    exit 1
  fi
fi

echo
echo "Continuing with:"
echo " - trust store file:        $trust_store_file"
echo " - trust store private key: $trust_store_private_key_file"

mkdir $KEYSTORE_WORKING_DIRECTORY

echo
echo "Now, a keystore will be generated. Each broker and logical client needs its own"
echo "keystore. This script will create only one keystore. Run this script multiple"
echo "times for multiple keystores."
echo
echo "You will be prompted for the following:"
echo " - A keystore password. Remember it."
echo " - Personal information, such as your name."
echo "     NOTE: currently in Kafka, the Common Name (CN) does not need to be the FQDN of"
echo "           this host. However, at some point, this may change. As such, make the CN"
echo "           the FQDN. Some operating systems call the CN prompt 'first / last name'"
echo " - A key password, for the key being generated within the keystore. Remember this."

# To learn more about CNs and FQDNs, read:
# https://docs.oracle.com/javase/7/docs/api/javax/net/ssl/X509ExtendedTrustManager.html

docker run --rm -it -v "$(pwd):/workspace" --workdir /workspace openjdk:latest keytool \
  -noprompt -alias localhost -keystore $KEYSTORE_WORKING_DIRECTORY/$KEYSTORE_FILENAME \
  -validity $VALIDITY_IN_DAYS -genkey -keyalg RSA -storepass $PASSWORD -keypass $PASSWORD

echo
echo "'$KEYSTORE_WORKING_DIRECTORY/$KEYSTORE_FILENAME' now contains a key pair and a"
echo "self-signed certificate. Again, this keystore can only be used for one broker or"
echo "one logical client. Other brokers or clients need to generate their own keystores."

echo
echo "Fetching the certificate from the trust store and storing in $CA_CERT_FILE."
echo
echo "You will be prompted for the trust store's password (labeled 'keystore')"

docker run --rm -it -v "$(pwd):/workspace" --workdir /workspace openjdk:latest keytool \
 -keystore $trust_store_file -export -alias CARoot -rfc -file $CA_CERT_FILE

echo
echo "Now a certificate signing request will be made to the keystore."
echo
echo "You will be prompted for the keystore's password."
docker run --rm -it -v "$(pwd):/workspace" --workdir /workspace openjdk:latest keytool \
  -keystore $KEYSTORE_WORKING_DIRECTORY/$KEYSTORE_FILENAME -alias localhost \
  -certreq -file $KEYSTORE_SIGN_REQUEST

echo
echo "Now the trust store's private key (CA) will sign the keystore's certificate."
echo
echo "You will be prompted for the trust store's private key password."
openssl x509 -req -CA $CA_CERT_FILE -CAkey $trust_store_private_key_file \
  -in $KEYSTORE_SIGN_REQUEST -out $KEYSTORE_SIGNED_CERT \
  -days $VALIDITY_IN_DAYS -CAcreateserial
# creates $KEYSTORE_SIGN_REQUEST_SRL which is never used or needed.

echo
echo "Now the CA will be imported into the keystore."
echo
echo "You will be prompted for the keystore's password and a confirmation that you want to"
echo "import the certificate."
docker run --rm -it -v "$(pwd):/workspace" --workdir /workspace openjdk:latest keytool \
  -keystore $KEYSTORE_WORKING_DIRECTORY/$KEYSTORE_FILENAME -alias CARoot \
  -import -file $CA_CERT_FILE
rm $CA_CERT_FILE # delete the trust store cert because it's stored in the trust store.

echo
echo "Now the keystore's signed certificate will be imported back into the keystore."
echo
echo "You will be prompted for the keystore's password."
docker run --rm -it -v "$(pwd):/workspace" --workdir /workspace openjdk:latest keytool \
  -keystore $KEYSTORE_WORKING_DIRECTORY/$KEYSTORE_FILENAME -alias localhost -import \
  -file $KEYSTORE_SIGNED_CERT

echo
echo "All done!"
echo
echo "Delete intermediate files? They are:"
echo " - '$KEYSTORE_SIGN_REQUEST_SRL': CA serial number"
echo " - '$KEYSTORE_SIGN_REQUEST': the keystore's certificate signing request"
echo "   (that was fulfilled)"
echo " - '$KEYSTORE_SIGNED_CERT': the keystore's certificate, signed by the CA, and stored back"
echo "    into the keystore"
echo -n "Delete? [yn] "
read delete_intermediate_files

if [ "$delete_intermediate_files" == "y" ]; then
  rm $KEYSTORE_SIGN_REQUEST_SRL
  rm $KEYSTORE_SIGN_REQUEST
  rm $KEYSTORE_SIGNED_CERT
fi