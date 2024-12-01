import { inject } from '@angular/core';
import { Router } from '@angular/router';

import { StateService } from '../services/state/state.service';

export const authGuard = () => {
  const state = inject(StateService);
  const router = inject(Router);

  if (state.getAuthCookieToken()) {
    return true;
  }

  // Redirect to the login page
  return router.parseUrl('login');
};