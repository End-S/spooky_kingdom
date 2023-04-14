import { HTTP } from "@/common/http";
import type { ApiResponse } from "@/common/models/api.model";
import type { Authentication } from "@/common/models/auth.model";
import type { JwtPayload } from "jwt-decode";
import jwtDecode from "jwt-decode";
import { DateTime } from "luxon";

export type LoginResponse = ApiResponse<{ jwt: string }>;
export const login = async (auth: Authentication): Promise<LoginResponse> =>
  HTTP.post("login", {
    username: auth.username,
    password: auth.password,
  });

export const isAuthorised = (): boolean => {
  const jwt: string | null = localStorage.getItem("JWT");
  if (!jwt) return false;

  const decode: JwtPayload = jwtDecode(jwt);
  if (!decode.exp) return false;

  const expiry: number = decode.exp;
  const now: number = DateTime.now().toSeconds();

  return now < expiry;
};
