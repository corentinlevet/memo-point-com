<script setup lang="ts">
  import goServer from '@/api/go-server';

  import '@/assets/Auth/style.css';
</script>

<template>
  <div class="page-container">
    <div class="auth-container">
      <div class="auth-header">
        <div id="auth-title" class="text">Sign up</div>
        <div class="underline"></div>
      </div>
      <div class="inputs">
        <div class="input">
          <img src="@/assets/Auth/person.png" alt="user icon" />
          <input type="text" v-model="username" placeholder="Username" required />
        </div>

        <div id="email-input" style="display: flex;" class="input">
          <img src="@/assets/Auth/email.png" alt="email icon" />
          <input type="email" v-model="email" placeholder="Email" required />
        </div>

        <div class="input">
          <img src="@/assets/Auth/password.png" alt="password icon" />
          <input type="password" v-model="password" placeholder="Password" required />
        </div>
      </div>

      <div class="submit-container">
        <div id="btn-sign-up" class="submit" @click="handleSignUp">Sign up</div>
        <div id="btn-log-in" class="submit gray" @click="handleLogIn">Log in</div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
  let username = '';
  let email = '';
  let password = '';
  let isOnSignUp = true;

  async function logIn() {
    goServer.logIn(username, password).then((response) => {
      if (response.status === 200)
      {
        const auth_token = response.data.auth_token;
        goServer.setAuthToken(auth_token);

        console.log('Logged in');
      }
      else
      {
        console.log('Unknown error');
      }
    }).catch((error) => {
      if (error.response.status === 404)
      {
        console.log(error.response.data.message);
      }
      else
      {
        console.log('Server error');
      }
    });
  }

  async function signUp() {
    goServer.signUp(username, email, password).then((response) => {
      if (response.status === 200)
      {
        const auth_token = response.data.auth_token;
        goServer.setAuthToken(auth_token);

        console.log('Signed up');
      }
      else
      {
        console.log('Unknown response');
      }
    }).catch((error) => {
      if (error.response.status === 400)
      {
        console.log("Bad request");
      }
      else
      {
        console.log('Server error');
      }
    });
  }

  function handleLogIn() {
    if (isOnSignUp)
    {
      isOnSignUp = false;
      const authTitle = document.getElementById('auth-title');
      if (authTitle)
      {
        authTitle.innerHTML = 'Log in';
      }
      const btnSignUp = document.getElementById('btn-sign-up');
      if (btnSignUp)
      {
        btnSignUp.classList.add('gray');
      }
      const btnLogIn = document.getElementById('btn-log-in');
      if (btnLogIn)
      {
        btnLogIn.classList.remove('gray');
      }
      const emailInput = document.getElementById('email-input');
      if (emailInput)
      {
        emailInput.style.display = 'none';
      }
    }
    else
    {
      logIn();
    }
  }

  function handleSignUp() {
    if (!isOnSignUp)
    {
      isOnSignUp = true;
      const authTitle = document.getElementById('auth-title');
      if (authTitle)
      {
        authTitle.innerHTML = 'Sign up';
      }
      const btnSignUp = document.getElementById('btn-sign-up');
      if (btnSignUp)
      {
        btnSignUp.classList.remove('gray');
      }
      const btnLogIn = document.getElementById('btn-log-in');
      if (btnLogIn)
      {
        btnLogIn.classList.add('gray');
      }
      const emailInput = document.getElementById('email-input');
      if (emailInput)
      {
        emailInput.style.display = 'flex';
      }
    }
    else
    {
      signUp();
    }
  }
</script>
