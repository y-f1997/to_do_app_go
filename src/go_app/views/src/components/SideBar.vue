<template>
<div class="container-fluid">
  <div style="margin-top: 100px;">
    <ul class="list-group">
      <router-link :to="{name:'Home'}" active-class="active--link" exact>
        <li class="list-group-item">
          <font-awesome-icon icon="home"></font-awesome-icon> Home
        </li>
      </router-link>

      <!-- <router-link :to="{name:'WhatDone', params:{action:'post'}}">
        <li class="list-group-item">Daily Done</li>
      </router-link> -->

      <router-link :to="{name:'TodoList'}" active-class="active--link" exact>
        <li class="list-group-item">To Do</li>
      </router-link>
      
    </ul>
  </div>
</div>
</template>
<style>
.active--link li {
  color: black;
  background-color: #cce5ff;
}
</style>

