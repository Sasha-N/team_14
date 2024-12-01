import { Routes } from '@angular/router';
import { LoginPageComponent } from './components/login-page/login-page.component';
import { MainPageComponent } from './components/main-page/main-page.component';

export const routes: Routes = [{
    path: 'login',
    component: LoginPageComponent
  },
  {
    path: '',
    component: MainPageComponent
  }];
