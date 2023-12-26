<script setup lang="ts">
  import goServer from '@/api/go-server';

  import '@/assets/Auth/style.css';
</script>

<template>
  <div class="page-container">
    <div class="auth-container">
      <div class="auth-header">
        <div id="auth-title" class="text">{{ currentStatus === 'sign-up' ? 'Sign up' : 'Log in' }}</div>
        <div class="underline"></div>
      </div>

      <div class="inputs">
        <div class="input">
          <img src="@/assets/Auth/person.png" alt="user icon" />
          <input type="text" v-model="username" placeholder="Username" required />
        </div>

        <div class="input" v-if="currentStatus === 'sign-up'">
          <img src="@/assets/Auth/email.png" alt="email icon" />
          <input type="email" v-model="email" placeholder="Email" required />
        </div>

        <div class="input">
          <img src="@/assets/Auth/password.png" alt="password icon" />
          <input type="password" v-model="password" placeholder="Password" required />
        </div>
      </div>

      <div class="submit-container">
        <div id="btn-sign-up" class="submit" v-on:click="handleSignUp">Sign up</div>
        <div id="btn-log-in" class="submit gray" v-on:click="handleLogIn">Log in</div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
  export default {
    data() {
      return {
        username: '',
        email: '',
        password: '',
        currentStatus: 'sign-up',
      };
    },
    methods: {
      auth() {
        if (!this.validateFields()) {
          alert('Veuillez remplir tous les champs.');
          return;
        }

        if (this.currentStatus === 'login') {
          goServer.logIn(this.username, this.password).then((response) => {
            const authToken = response.data.auth_token;
            goServer.setAuthToken(authToken);
            this.$router.push('/home');
          }).catch((error) => {
            console.error(error);
          });
        } else {
          goServer.signUp(this.username, this.email, this.password).then((response) => {
            const authToken = response.data.auth_token;
            goServer.setAuthToken(authToken);
            this.$router.push('/home');
          }).catch((error) => {
            console.error(error);
          });
        }
      },
      handleSignUp() {
        if (this.currentStatus === 'login') {
          document.getElementById('btn-sign-up')?.classList.remove('gray');
          document.getElementById('btn-log-in')?.classList.add('gray');
          this.currentStatus = 'sign-up';

          return;
        }

        if (this.currentStatus === 'sign-up') {
          this.auth();
        }
      },
      handleLogIn() {
        if (this.currentStatus === 'sign-up') {
          document.getElementById('btn-log-in')?.classList.remove('gray');
          document.getElementById('btn-sign-up')?.classList.add('gray');
          this.currentStatus = 'login';

          return;
        }

        if (this.currentStatus === 'login') {
          this.auth();
        }
      },
      validateFields() {
        if (this.currentStatus === 'sign-up') {
          return this.username && this.email && this.password;
        }

        if (this.currentStatus === 'login') {
          return this.username && this.password;
        }

        return false;
      },
    },
  };
</script>
