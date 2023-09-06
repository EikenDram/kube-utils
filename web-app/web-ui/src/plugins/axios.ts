import axios, { AxiosInstance } from "axios";
import { inject } from "vue";
import type { App } from "vue";

// vue injection key
const injectionKey = Symbol("api");

// Keycloak API class
class Keycloak {
  http: AxiosInstance;

  // get all logs
  getAll(): Promise<any> {
    return this.http.get("/keycloak/list");
  }

  // get log content
  get(uid: string): Promise<any> {
    return this.http.get(`/keycloak/log/${uid}`);
  }

  // upload file
  upload(file: File): Promise<any> {
    const data = new FormData();
    data.append("file", file);
    return this.http.post("/keycloak/upload", data, {
      headers: {
        "Content-Type": "multipart/form-data",
      },
    });
  }

  // preview file format
  preview(data: any): Promise<any> {
    return this.http.post("/keycloak/preview", data);
  }

  // process file
  process(data: any): Promise<any> {
    return this.http.post("/keycloak/process", data);
  }

  constructor(http: AxiosInstance) {
    this.http = http;
  }
}

// API class
class API {
  http: AxiosInstance;

  Keycloak: Keycloak;

  constructor() {
    this.http = axios.create({
      baseURL: import.meta.env.VITE_API,
      withCredentials: false,
    });
    this.Keycloak = new Keycloak(this.http);
  }
}

export const useAPI = () => inject(injectionKey) as API;

export default {
  install: (app: App) => {
    const http = new API();
    app.provide(injectionKey, http);
    app.config.globalProperties.$http = http;
  },
};
