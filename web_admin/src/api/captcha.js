import request from '@/utils/request'

export function captcha(data) {
  return request({
    url: '/v1/captcha',
    method: 'get'
  })
}
