import axios from 'axios';

const HTTP = axios.create({
  baseURL: 'http://localhost:8585/api/',
  headers: {
    Authorization: 'Bearer {token}',
  },
});

export { HTTP };
