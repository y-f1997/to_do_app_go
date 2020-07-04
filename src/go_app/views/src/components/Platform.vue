<template>
<div>
  <slot></slot>
  <div class="container-fluid" id="platform">
    <div class="row" style="min-height: 1000px;">
      <div class="col-md-2 pl-0 pr-0 bg-light">
        <SideBar></SideBar>
      </div>
      <div class="col-md-10" style="margin-top: 100px;">
        <router-view></router-view>
        <Error4xx></Error4xx>
      </div>
    </div>
  
  </div>

</div>
</template>

<script>
import Error4xx from './Error4xx.vue';
import SideBar from './SideBar.vue';
export default {
  components: {
    Error4xx,
    SideBar
  }
}
</script>
