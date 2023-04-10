import axios from "axios";
import type { ErrorResponse } from "./models/api.model";

function baseURL() {
  if (process.env.NODE_ENV === "development") {
    return `${window.location.protocol}//${window.location.hostname}:8585/api`;
  }
  return `${window.location.origin}/api/`;
}

const HTTP = axios.create({
  baseURL: baseURL(),
});

// intercept all responses from the server
HTTP.interceptors.response.use(
  (res) => {
    // happy path, return data sent from the server
    return res.data;
  },
  (err): ErrorResponse => {
    let errorRes = { error: "Server Error", status: "500" };
    // unhappy path, manage and return an error response
    if (err?.response?.data?.error) {
      errorRes = {
        error: err.response.data.error,
        status: err?.response?.status ?? "500",
      };
    }
    return errorRes;
  }
);

export { HTTP };
