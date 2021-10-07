import axios from 'axios';

const HTTP = axios.create({
  baseURL: 'http://192.168.1.248:8585/api/',
  // baseURL: 'http://localhost:8585/api/',
  headers: {
    Authorization: 'Bearer {token}',
  },
});

export { HTTP };
