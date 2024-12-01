import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

export interface UserLoginApiResponse {
  token: string
}

export interface UserRegistrApiResponse {
  message: string,
}

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  constructor(private _httpClient: HttpClient) { }

  public login(authData: { name: string, password: string }): Observable<UserLoginApiResponse> {
    const apiUrl = `http://localhost:8080/users/login`;

    const headers = new HttpHeaders({
      'Content-Type': 'application/json',
    });

    const body = authData;

    return this._httpClient.post<UserLoginApiResponse>(apiUrl, body, {
      headers: headers,
      withCredentials: false
    });
  }

  public registr(authData: { name: string, password: string }): Observable<UserRegistrApiResponse> {
    const apiUrl = `http://localhost:8080/users`;

    const headers = new HttpHeaders({
      'Content-Type': 'application/json',
    });

    const body = authData;

    return this._httpClient.post<UserRegistrApiResponse>(apiUrl, body, {
      headers: headers,
      withCredentials: false
    });
  }
}
