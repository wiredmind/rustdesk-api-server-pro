import { request } from '../request';

export function fetchDevicesList(params: any) {
  return request<Api.Devices.DevicesList>({ url: '/devices/list', params });
}

export function deleteDevices(data: { ids: number[] }) {
  return request({ url: '/devices/delete', method: 'post', data });
}
