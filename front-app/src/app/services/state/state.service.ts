import { Injectable } from '@angular/core';
import { Observable, tap, shareReplay, BehaviorSubject } from 'rxjs';
import { UserLoginApiPayload } from '../../components/login-page/login-page.component';
import Cookies from 'js-cookie';
import { ApiService } from '../api/api.service';

@Injectable({
  providedIn: 'root'
})
export class StateService {

  private _user = {
    data: new BehaviorSubject<UserLoginApiPayload | null>(null),
    loading: new BehaviorSubject<boolean>(false),
    error: new BehaviorSubject<Error | null>(null),
  };

  constructor(private apiService: ApiService) { }

  public setAuthCookieToken(token: string): void {
    Cookies.set('token', token, { sameSite: 'strict', path: '/', secure: true, expires: 30 })
  }

  
}
