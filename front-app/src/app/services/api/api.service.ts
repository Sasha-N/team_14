import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

export interface UserApiResponse {
  id: number,
  email: string,
  name: string | null,
  token: string
}

export interface UserLoginApiResponse {
  user: UserApiResponse,
  token: string,
}

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  constructor(private _httpClient: HttpClient) { }

  public login(authData: { login: string, password: string }): Observable<UserLoginApiResponse> {
    const apiUrl = `http://localhost:3000/auth/login`;

    const headers = new HttpHeaders({
      'Content-Type': 'application/json',
    });

    const body = authData;

    return this._httpClient.post<UserLoginApiResponse>(apiUrl, body, {
      headers: headers,
      withCredentials: false
    });
  }
}
