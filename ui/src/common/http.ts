import axios from 'axios';

function baseURL() {
  if (process.env.NODE_ENV === 'development') {
    return `${window.location.protocol}//${window.location.hostname}:8585/api`;
  }
  return `${window.location.origin}/api/`;
}

const HTTP = axios.create({
  // baseURL: 'http://192.168.1.248:8585/api/',
  // `${ window.location.protocol }//${ window.location.hostname }:${ this.environment.port }/api`;
  // baseURL: 'http://localhost:8585/api/',
  // baseURL: `${ window.location.origin }/api/`,
  baseURL: baseURL(),
  headers: {
    Authorization: 'Bearer {token}',
  },
});

export { HTTP };
