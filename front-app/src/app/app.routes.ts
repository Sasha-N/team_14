import { Routes } from '@angular/router';
import { LoginPageComponent } from './components/login-page/login-page.component';
import { MainPageComponent } from './components/main-page/main-page.component';
import { RegistrationPageComponent } from './components/registration-page/registration-page.component';
import { authGuard } from './auth/auth.guard';

export const routes: Routes = [{
  path: 'login',
  component: LoginPageComponent
},
{
  path: '',
  component: MainPageComponent,
  canActivate: [authGuard]
},
{
  path: 'registration',
  component: RegistrationPageComponent,
}];
