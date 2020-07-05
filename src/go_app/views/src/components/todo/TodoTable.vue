<template>
<div>
  <table class="table table-hover list">
    <thead :class="todo_status">
      <th style="width:20%;" v-if="todo_status === 'todo'">期限</th>
      <th style="width:20%;" v-if="todo_status === 'fin'">終了日時</th>
      <th style="width:10%;">重要度</th>
      <th style="width:10%;">カテゴリ</th>
      <th style="width:55%;">ToDo</th>
      <!-- <th>追加情報</th> -->
      <th style="width:5%;">削除</th>
      <th></th>
      <th style="display: none;">更新日時</th>
    </thead>
    <tbody>
      <tr v-for="todo in todo_list" v-bind:key="todo.crtTimestamp" style="cursor:pointer;">
        <td v-if="todo_status === 'todo'" @click="routeDetail(todo)">{{convertDateFormat(todo.timeToBeDone)}}</td>
        <td v-if="todo_status === 'fin'" @click="routeDetail(todo)">{{convertDateFormat(todo.timeDone)}}</td>
        <td :class="importanceColor(todo.importance)" @click="routeDetail(todo)">{{todo.todoImportanceRes}}</td>
        <td @click="routeDetail(todo)">{{todo.todoCategoryRes}}</td>
        <td @click="routeDetail(todo)">{{todo.toDo}}</td>
        <!-- <td>{{todo.additionalInfo}}</td> -->
        <td>
          <font-awesome-icon icon="minus-circle" style="color: red;" @click="deleteTodo(todo)" />
        </td>
        <td style="display: none">{{todo.crtTimestamp}}</td>
      </tr>
    </tbody>
  </table>

</div>
</template>

<script>
import axios from 'axios';
export default {
  props: ['todo_list', 'todo_status'],
  data(){
    return {
      urlDeleteTodo: "/deleteTodo",
    }
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
    routeDetail: function (todo) {
      var lastWhatDone = {};
      if (this.todo_list.length>0){
        lastWhatDone = this.todo_list[this.todo_list.length - 1];
      }else{
        lastWhatDone.endDateTime = null;
      }
      
      var to = {
        //新規作成用URL
        name: 'Todo',
        params: {
          action: 'post',
          postStartDateTime: lastWhatDone.endDateTime
        }
      }
      if (todo !== null) {
        //更新用URL(crtTimestampでサーチ)
        to.params.action = todo.crtTimestamp;
      }
      this.$router.push(to)
    },
    deleteTodo(doc) {
      if (!confirm("削除しますか？")) {
        return;
      }
      var self = this;
      axios.post(self.urlDeleteTodo, doc).then(function (res) {
        console.log(res);
        self.$emit('parent-get-list')
      });
    },
    importanceColor: function(importance){
      if (importance === "1"){
        return 'bg-light-red';
      }else if (importance === "2"){
        return 'bg-light-yellow'
      }else{
        return 'bg-light-green'
      }
    }
  }
  
}
</script>

<style>
  table .todo {
    background: #b8b8b7;
    /* background: #f8d7da; */
  }

  table .fin {
    background: #2bf502;
  }
    table{
    margin-top: 1em;
  }

  .list {
    table-layout: fixed;
  }
</style>