import { AfterViewInit, ChangeDetectorRef, Component, OnDestroy } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { Subject, combineLatest, takeUntil } from 'rxjs';
import { StateService } from '../../services/state/state.service';
import { UserLoginApiResponse } from '../../services/api/api.service';

@Component({
  selector: 'app-registration-page',
  imports: [
    MatInputModule,
    ReactiveFormsModule,
    MatIconModule,
    MatButtonModule,
    MatFormFieldModule
  ],
  templateUrl: './registration-page.component.html',
  styleUrl: './registration-page.component.css'
})
export class RegistrationPageComponent implements OnDestroy, AfterViewInit{
  public loginForm = new FormGroup({
    login: new FormControl<string>('', {
      validators: [Validators.required],
      nonNullable: true
    }),
    password: new FormControl<string>('', {
      validators: [Validators.required],
      nonNullable: true
    }),
    confirmPassword: new FormControl<string>('', {
      validators: [Validators.required],
      nonNullable: true
    })
  });

  public userData: { data: UserLoginApiResponse | null, error: Error | null, loading: boolean } = {
    data: null,
    error: null,
    loading: false,
  };

  private _destroyed = new Subject<boolean>();

  constructor(
    private _state: StateService,
    private _router: Router,
    private _changeDetectorRef: ChangeDetectorRef
  ) { }

  public ngOnDestroy(): void {
    this._destroyed.next(true);
    this._destroyed.complete();
  }

  public ngAfterViewInit(): void {
    combineLatest({
      ...this._state.user,
    }).pipe(
      takeUntil(this._destroyed),
    ).subscribe(({ data, error, loading }) => {
      this.userData.data = data;
      this.userData.error = error;
      this.userData.loading = loading;

      if (loading) {
        this.loginForm.disable();
      } else {
        this.loginForm.enable();
      }

      this._changeDetectorRef.markForCheck();
    })
  }


  public login(): void {
    this._state.login({
      login: this.loginForm.controls.login.value,
      password: this.loginForm.controls.password.value,
    }).pipe(
      takeUntil(this._destroyed),
    ).subscribe(() => {
      this._router.navigate(['']);
    })
  }
}
