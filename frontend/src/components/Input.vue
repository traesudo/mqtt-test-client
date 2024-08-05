<template>
  <div class="config">
  <el-form :model="form" label-width="auto" style="max-width: 600px">
    <el-form-item label="Host">
      <el-input v-model="form.host" />
    </el-form-item>
    <el-form-item label="Port">
      <el-input v-model="form.port" />
    </el-form-item>
    <el-form-item label="Message">
      <el-input v-model="form.message" />
    </el-form-item>
    <el-form-item label="Topic">
      <el-input v-model="form.topic" />
    </el-form-item>
    <el-form-item label="Username">
      <el-input v-model="form.username" />
    </el-form-item>
    <el-form-item label="Password">
      <el-input v-model="form.password" />
    </el-form-item>
    <el-form-item label="结束时间">
      <el-col :span="11">
        <el-date-picker
          v-model="form.date1"
          type="date"
          placeholder="Pick a date"
          style="width: 100%"
        />
      </el-col>
      <el-col :span="2" class="text-center">
        <span class="text-gray-500">-</span>
      </el-col>
      <el-col :span="11">
        <el-time-picker
          v-model="form.date2"
          placeholder="Pick a time"
          style="width: 100%"
        />
      </el-col>
    </el-form-item>
    <div class="slider-demo-block">
      <span class="demonstration">并发数量</span>
      <el-slider v-model="form.concurrencies" />
    </div>
    <el-form-item label="并发数量">
      <el-select v-model="form.concurrencies" placeholder="选择你的并发数量">
        <el-option label="1" value="1" />
        <el-option label="100" value="100" />
        <el-option label="200" value="200" />
        <el-option label="500" value="500" />
        <el-option label="3000" value="3000" />
        <el-option label="5000" value="5000" />
      </el-select>
      <el-button type="primary" @click="onSubmit" style="margin: 10px;">Start</el-button>
      <el-button @click="stopHandle">Stop</el-button>
    </el-form-item>
  </el-form>
</div>
<div class="recommended">
  <el-button type="info" @click="useConfig" plain>推荐本地配置1</el-button>
</div>

</template>

<script lang="ts" setup>
import { reactive } from 'vue';
import { Start,Stop } from "../../wailsjs/go/main/App";
import { main } from '../../wailsjs/go/models';
import { ElMessageBox } from 'element-plus'
let form = reactive({
  host: '',
  username: '',
  password: '',
  name: '',
  date1: '',
  date2: '',
  concurrencies: 0,
  topic:'',
  message: '',
  port:'',
});
function useConfig() {
  form.concurrencies = 100
  form.host = '127.0.0.1'
  form.username = 'sender'
  form.password = 'sender'
  form.topic = 'Test11/te'
  form.name = '测试1'
  form.port = '1883'
}

const onSubmit = async () => {
  try {
    console.log('submit! check data form', form);
    console.log('submit! check data form', form.username);

    if (form.username.length<= 0){
      console.log("这里有问题吧")
        ElMessageBox.alert('username 必填呀 哥们', 'Title', {
      })
      return
    }
    if (form.password.length<= 0){
        ElMessageBox.alert('password 必填呀 哥们', 'Title', {
      })
      return
    }
    if (form.host.length<= 0){
        ElMessageBox.alert('host 必填呀 哥们', 'Title', {
      })
      return
    }
    if (form.concurrencies<= 0){
        ElMessageBox.alert('不是，并发数量选都不选是吧？', 'Title', {
      })
      return
    }

    let input = new main.StartInput();
    input.username = form.username;
    input.password = form.password;
    input.concurrencies = form.concurrencies;
    input.name = form.name;
    input.date1 = form.date1;
    input.date2 = form.date2;
    input.message = form.message;
    input.topic = form.topic;
    input.host = form.host;
    input.port = form.port;
    
    // Call Start and await its result
    const result = await Start(input);

    // Handle the result as needed
    if (result) {
      console.log('Start succeeded:', result);
    } else {
      console.error('Start failed:', result);
    }
  } catch (error) {
    console.error('Error calling Start:', error);
  }
};

const stopHandle = async()=>{
  try{
    let result  =await Stop()
    if (result) {
      console.log('Stop succeeded:', result);
    } else {
      console.error('Stop failed:', result);
    }    
  }catch(error){
    console.error('Error calling Stop:', error);
  }
}
</script>


<style>
.recommended{
  float: left;
  text-align: center;
  margin: 10px;
  padding: 10px;
  /* border: 1px solid #ccc; */
  border-radius: 5px;
  width: auto;
  height: 20%;
  /* line-height: 10px; */
}
.config{
  width: 80%;
  height: auto;
  margin: 10px;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 5px;
  float: left;
  text-align: center;

}
</style>