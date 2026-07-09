import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { AuthService } from '@arch-shared/data-source';

@Component({
  changeDetection: ChangeDetectionStrategy.OnPush,
  selector: 'lib-login-page',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss'],
  imports: [ReactiveFormsModule],
})
export class LoginComponent {

  //DI
  readonly fb = inject(FormBuilder);
  readonly authService = inject(AuthService);

  readonly loginForm = this.fb.group({
    username: ['', Validators.required],
    password: ['', Validators.required],
  });

  onSubmit() {
    if (this.loginForm.valid) {
      const { username, password } = this.loginForm.value;
      this.authService.login({
        username: username as string,
        password: password as string,
      }).subscribe(o => {
        console.log('Login successful', o);
      });
      // Handle login logic here
    } else {
      console.log('Form is invalid');
    }
  }

}
