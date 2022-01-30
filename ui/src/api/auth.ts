import { AxiosResponse } from 'axios';
import jwtDecode, { JwtPayload } from 'jwt-decode';
import dayjs from 'dayjs';
import { HTTP } from '@/common/http';
import { Authentication } from '@/common/models/auth.model';

type LoginResponse = Promise<AxiosResponse<{ jwtToken: string }>>
export const login = async (auth: Authentication): LoginResponse => HTTP
  .post('login', {
    username: auth.username,
    password: auth.password,
  });

export const isAuthorised = (): boolean => {
  const jwt: string | null = localStorage.getItem('JWT');
  if (!jwt) return false;

  const decode: JwtPayload = jwtDecode(jwt);
  if (!decode.exp) return false;

  const expiry: number = decode.exp;
  const now: number = dayjs().unix();

  return now < expiry;
};
