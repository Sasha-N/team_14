import { Injectable } from '@angular/core';
import { Observable, tap, shareReplay, BehaviorSubject } from 'rxjs';
import Cookies from 'js-cookie';
import { ApiService, UserLoginApiResponse } from '../api/api.service';

@Injectable({
  providedIn: 'root'
})
export class StateService {

  public user = {
    data: new BehaviorSubject<UserLoginApiResponse | null>(null),
    loading: new BehaviorSubject<boolean>(false),
    error: new BehaviorSubject<Error | null>(null),
  };

  constructor(private apiService: ApiService) { }

  public setAuthCookieToken(token: string): void {
    Cookies.set('token', token, { sameSite: 'strict', path: '/', secure: true, expires: 30 })
  }

  public login(authData: { login: string, password: string }): Observable<UserLoginApiResponse> {
    this.user.loading.next(true);
    this.user.error.next(null);

    return this.apiService.login(authData).pipe(
      tap({
        next: response => {
          this.setAuthCookieToken(response.token);
          this.user.data.next({ ...response });
          this.user.loading.next(false);
          this.user.error.next(null);
        },
        error: error => {
          this.user.data.next(null);
          this.user.loading.next(false);
          this.user.error.next(error);
        },
      }),
      shareReplay({ bufferSize: 1, refCount: false })
    );
  }
}
