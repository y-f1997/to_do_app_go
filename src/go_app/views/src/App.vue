<template>
<div id="app">
  <!-- 全URL共通部分 -->
  <Platform>
  <nav class="navbar fixed-top navbar-dark bg-dark">
    <div class="row">
      <div class="col-md-6">
        <div class="col-md-3">
          <router-link :to="{name: 'Home'}" class="navbar-brand">My Portal</router-link>
        </div>
      </div>
      <div class="col-md-6">
        <div class="float-right">
          <router-link to="/signUp">
            <button class="btn btn-outline-light">
              Sign Up
            </button>
          </router-link>
        </div>

      </div>
    </div>
  </nav>
  </Platform>

</div>
</template>
<script>
import Platform from './components/Platform.vue';
export default {
  components:{
    Platform
  }
}
</script>
<style>

  .bg-orange{
    background: #faa302;
  }

  .bg-green{
    background: #2bf502;
  }


  .bg-light-red{
    background: #f8d7da;
  }

  .bg-light-yellow{
    background: #fff3cd;
  }

  .bg-light-green {
    background-color: rgb(240, 255, 246);
  }

</style>
