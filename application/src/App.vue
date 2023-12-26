<template>
  <div>
    <Navbar v-if="!$route.name || $route.name !== 'auth'"/>

    <RouterView />
  </div>
</template>

<script setup lang="ts">
  import { RouterView, useRoute } from 'vue-router'
  import 'bootstrap/dist/css/bootstrap.min.css';

  import Navbar from '@/components/Navbar.vue';

  const $route = useRoute();
</script>
