<script setup lang="ts">
  import '@/assets/Navbar/style.css';
</script>

<template>
  <header id="navHeader">
    <h3>Memo-point-com</h3>

    <nav id="navBar" ref="navRef">
      <router-link to="/home">Home</router-link>
      <router-link to="/memos">Memos</router-link>
      <router-link to="/about">About</router-link>
      <button @click="showNavbar" class="nav-btn nav-close-btn">
        <font-awesome-icon icon="fa-times" />
      </button>
      <button @click="logOut" class="btn btn-danger">Log out</button>
    </nav>
    <button @click="showNavbar" class="nav-btn">
      <font-awesome-icon icon="fa-bars" />
    </button>
  </header>
</template>

<script lang="ts">
  import { ref } from 'vue';
  import goServer from '@/api/go-server';

  const navRef = ref(null);

  const showNavbar = () => {
    navRef.value.classList.toggle("responsive-nav");
  }

  export default {
    methods: {
      logOut() {
        this.$router.push('/');
        goServer.removeAuthToken();
      }
    }
  }
</script>
