<template>
<div>
  <div class="container-fluid">
    <div class="row">
      <div class="col-md-11">
        <h4 class="form-inline">
          <v-date-picker :mode="'range'" v-model="selectedDate" @input="getList()" :format="'YYYY-MM-DDTHH:mm:ss.sssZ'" style="width:270px;"></v-date-picker>
        </h4>
      </div>
      <div class="col-md-1">
      <h2 class="float-right">
        <div @click="routeDetail(null)" style="color:blue; cursor:pointer">
          <font-awesome-icon icon="file-alt" />
        </div>
      </h2>

      </div>
    </div>
    <div class="row">
      <table class="table table-hover list">
        <thead class="thead-dark">
          <th style="width:15%;">開始日時</th>
          <th style="width:15%;">終了日時</th>
          <th style="width:10%;">カテゴリ</th>
          <th style="width:50%;">何をしたか</th>
          <!-- <th>追加情報</th> -->
          <th style="width:10%;">削除</th>
          <th></th>
          <th style="display: none;">更新日時</th>
        </thead>
        <tbody>
        <tr v-for="whatDone in lists" v-bind:key="whatDone.crtTimestamp" style="cursor:pointer;">
          <td @click="routeDetail(whatDone)">{{convertDateFormat(whatDone.startDateTime)}}</td>
          <td @click="routeDetail(whatDone)">{{convertDateFormat(whatDone.endDateTime)}}</td>
          <td @click="routeDetail(whatDone)">{{whatDone.categoryRes}}</td>
          <td @click="routeDetail(whatDone)">{{whatDone.whatDone}}</td>
          <!-- <td>{{whatDone.additionalInfo}}</td> -->
          <td>
            <font-awesome-icon icon="minus-circle" style="color: red;" @click="deleteWhatDone(whatDone)" />
          </td>
          <td style="display: none">{{whatDone.crtTimestamp}}</td>
        </tr>
        </tbody>
      </table>
    </div>

  </div>
</div>
</template>

<script>
import axios from 'axios';
export default {
  name: 'Home',
  data() {
    return {
      urlGetList: "/getList",
      urlDeleteWhatDone: "/deleteWhatDoneDetail",
      lists: [],
      selectedDate: {
        start:new Date(),
        end:new Date()
        }
    }
  },
  mounted: function () {
    this.getList()
  },
  methods: {
    convertDateFormat: function (date) {
      if (date === null){
        return null;
      }
      var formattingDate = new Date(date);
      var year = formattingDate.getFullYear();
      var month = ("0" + (formattingDate.getMonth() + 1)).slice(-2);
      var day = ("0" + formattingDate.getDate()).slice(-2);
      var hour = formattingDate.getHours();
      var min = formattingDate.getMinutes();
      var newFormatDate = year + "/" + month + "/" + day + " " + hour + "時" + min + "分";
      return newFormatDate;
    },
    routeDetail: function (whatDone) {
      var lastWhatDone = {};
      if (this.lists.length>0){
        lastWhatDone = this.lists[this.lists.length - 1];
      }else{
        lastWhatDone.endDateTime = null;
      }
      
      var to = {
        //新規作成用URL
        name: 'WhatDone',
        params: {
          action: 'post',
          postStartDateTime: lastWhatDone.endDateTime
        }
      }
      if (whatDone !== null) {
        //更新用URL
        to.params.action = whatDone.crtTimestamp;
      }
      this.$router.push(to)
    },
    getList() {
      var self = this;
      var queryParams = {
        params:{
          spanStart: self.selectedDate.start,
          spanEnd: self.selectedDate.end
        }
      }
      console.log(queryParams)
      axios.get(this.urlGetList,queryParams).then(function (res) {
        var list = res.data.list;
        if (list === null) {
          list = [];
        }
        self.lists = list;
      })
    },
    deleteWhatDone(doc) {
      if (!confirm("削除しますか？")) {
        return;
      }
      var self = this;
      axios.post(self.urlDeleteWhatDone, doc).then(function (res) {
        console.log(res);
        self.getList();
      });
    }
  }

}
</script>

<style>
  table{
    margin-top: 1em;
  }
  .list {
    table-layout: fixed;
  }
</style>