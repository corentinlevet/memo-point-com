<script setup lang="ts">
  import goServer from '../api/go-server';

  import 'bootstrap/dist/css/bootstrap.min.css';
</script>

<template>
  <div class="page-container">
    <div class="login-container">
      <div class="login-header">
        <div class="text">Login</div>
        <div class="underline"></div>
      </div>
      <div class="inputs">
        <div class="input">
          <img src="../assets/Auth/person.png" alt="user icon" />
          <input type="text" v-model="username" placeholder="Username" required />
        </div>

        <div class="input">
          <img src="../assets/Auth/password.png" alt="password icon" />
          <input type="password" v-model="password" placeholder="Password" required />
        </div>
      </div>

      <div class="submit-container">
        <div class="submit" @click="submitForm">Login</div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
  let username = '';
  let password = '';

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

  const submitForm = () => {
    logIn();
  };
</script>

<style scoped>
  .page-container {
    margin: 0;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    height: 100vh;
    overflow: hidden;
    background: linear-gradient(#ca1111, #d17f03);
    display: flex;
    justify-content: center;
  }

  .login-container {
    display: flex;
    flex-direction: column;
    margin: auto;
    background: #fff;
    width: 600px;
    border-radius: 20px;
    min-height: 600px;
  }

  .login-header {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 10px;
    width: 100%;
    margin-top: 30px;
  }

  .text {
    color: #ca1212;
    font-size: 2rem;
    font-weight: 700;
  }

  .underline {
    width: 60px;
    height: 5px;
    background: #ca1212;
    border-radius: 10px;
  }

  .inputs {
    margin-top: 55px;
    display: flex;
    flex-direction: column;
    gap: 25px;
  }

  .input {
    display: flex;
    align-items: center;
    margin: auto;
    width: 480px;
    height: 80px;
    background: #EAEAEA;
    border-radius: 5px;
  }

  .input img {
    margin: 0px 30px;
  }

  .input input {
    height: 50px;
    width: 400px;
    background: transparent;
    border: none;
    outline: none;
    color: #797979;
    font-size: 20px;
  }

  .submit-container {
    display: flex;
    gap: 30px;
    margin: 20px auto;
    margin-top: auto;
    margin-bottom: 30px;
  }

  .submit {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 220px;
    height: 60px;
    color: #fff;
    background: #f89e17;
    border-radius: 50px;
    font-size: 20px;
    font-weight: 700;
    cursor: pointer;
  }

</style>