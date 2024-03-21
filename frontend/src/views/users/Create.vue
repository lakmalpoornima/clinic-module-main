<template>
<!-- <CContainer > -->
  
<!-- =============================================================Get User Input===================================================== -->
  <!-- <div  v-else>  -->
    <CRow>
      <CCol sm="6">
        <CCard>
          <CCardHeader>
            <div  v-if='this.id'>
            <strong>Update</strong><small>Form</small>
            </div>
            <div v-else>
              <strong>Create New user</strong><small>Form</small>
            </div>
            <!-- <CInput type="text" aria-placeholder="Search"/> -->
            <!-- {{Item}} -->
          </CCardHeader>
          <form  v-on:submit.prevent="Post_data_for_user" enctype="multipart/form-data" >
          <CCardBody>
            
              <CRow >
                <CCol sm="6">
                  <CInput  placeholder=" Username" v-model="posts.user_name" />
                
                  <CInput  placeholder="User ID" v-model="posts.userid" />
                  
                  <CInput  placeholder="E-Mail" v-model="posts.email" type="email"/>
                 
                  <CInput placeholder="User code" v-model="posts.user_code" type="text"/>
                   
                  <CInput placeholder="Password"  id="password" type="password"  v-model="posts.password_hash">
                    
                  <template #append-content><i class="font-weight-bold p-20" @click="show_password" ><CIcon name="cil-sun"  id="showtext"/> </i></template>
                  </CInput>
                
                  
               
                  
                </CCol>

                <!-- <CCol sm="4">
                   <div class="profiler">
                   
                      <cropper :src="imageData"  
                       
                       :stencil-props="{
                          handlers: {},
                          movable: false,
                          scalable: false,                     
                        }"
                        
                        :resize-image="{
                          adjustStencil: false
                        }"

                        :stencil-size="{
                          width: 200,
                          height: 200
                        }"
                        image-restriction="stencil"
                        

                        @change="change" />
                     <div class="base-image-input" :style="{ 'background-image': `url(${imageData})` }" @click="chooseImage" >
                      

                         <span  v-if="!imageData" class="placeholder" > Choose an Image </span>
                        <input class="file-input" ref="fileInput" type="file" @input="onSelectFile"   accept=".png, .jpg, .jpeg">
                         <span>hello{{this.getlocal}}</span> -->
                      <!-- </div>
                 
                  </div>
                </CCol>  -->

                <CCol sm="6">
                  <CInput  placeholder="Login ID" v-model="posts.login_id" />
                </CCol> 
            </CRow>

            <CRow >
                 <CCol sm="6">
                    <CInput placeholder="Telephone No" type="text" v-model="posts.telephone_no"   /> 
                 </CCol>

                 <CCol sm="6">
                  <CInput placeholder="Mobile No" type="text" v-model="posts.mobile_no"/>
                 </CCol>
               
            </CRow>

            <CRow >
                <CCol sm="12">
                  
                    <c-textarea  placeholder="Address" rows="3" type="text" v-model="posts.address" />
                </CCol>
            </CRow>

            <CRow >
              <CCol sm="4">
                <select name="cars" id="cars" class="m-2"  v-model="posts.company_code">
              
                <option disabled selected > Company Code </option>
                <option value="C001">C001</option>
                </select>
             </CCol>
            
           
            <CCol sm="4">
                <!-- <select name="cars" id="cars" class="m-2" v-model="posts.department_code"> -->
                  <select name="cars" id="cars" class="m-2"  v-model="posts.department_code">
                 <option disabled selected> Department Code </option>
                 
                <option value="C001">D001</option>
                <option value="C001">D002</option>
                </select>
            </CCol>

           <CCol sm="4">
              <select name="cars" id="cars" class="m-2"  v-model="posts.is_activate">
                <!-- <option value="">--Please choose an option--</option> -->
              
	
                     <option disabled selected> User Status </option>
                    <option value="true"> 🟢 Active </option>
                    <option value="false"> 🔴 Inactive</option>
                    
              </select>

              
          </CCol>
          
      </CRow>

          </CCardBody>
          <CCardFooter>
            <div class="d-flex justify-content-center" v-if='this.id'>
              <CButton  type="submit" color="info" sm="6" class="font-weight-bold" ><CIcon name="cilCheck"/>&nbsp;Submit</CButton> 
              &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
              <CButton  type="Reset" color="info" @click="back" sm="8" class="font-weight-bold" ><CIcon name="cil-cursor"/>&nbsp;Back</CButton>
            </div>

            <div v-else> 
              <CButton  type="submit" color="info" sm="6" class="font-weight-bold" ><CIcon name="cilCheck"/>&nbsp;Submit</CButton> 
              &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
              <CButton  type="Reset" color="info" @click="back" sm="8" class="font-weight-bold" ><CIcon name="cil-cursor"/>&nbsp;Back</CButton>
            </div>
            
          </CCardFooter>
          </form>
        </CCard>
      </CCol>
    </CRow>
   
  <!-- </div> -->
   <!-- </CContainer> -->
</template>



<script>

import * as Vue from 'vue' 
import axios from 'axios'
import VueAxios from 'vue-axios'; 

import useVuelidate from '@vuelidate/core';
import { Cropper } from 'vue-advanced-cropper';
import 'vue-advanced-cropper/dist/style.css';

import { required,email, integer} from '@vuelidate/validators';
// import fileupload from 'public/img/avatars';
import { reactive,computed} from 'vue';

export default{
  name:"Create",//Component Name
   inject: ["api"],
  // setup() {
  //   const state = reactive({
  //     userid:'',
  //     email:'',
  //     name:'',
  //   })
  // },
   

  data(){ //data part
  // console.log("Goo",imageData)
    return{
       imageData: null,
       progress :0,
      //  getlocal: null,
     
      posts:{
          userid:"",
          email:"",
          image: this.getlocal,
          telephone_no:"",
          mobile_no:"",
          address:"",
          company_code:"",
          department_code:"",
          is_activate:"",
          user_code:"",                                                             
          user_name:"",                                   
          login_id:"",                                   
          password_hash:"",  
      },
      
      coordinates: {
            width: 0,
            height: 0,
            left: 0,
            top: 0,
          },
    }
   
  },
  props:{
    id:{ type:integer, required:true}
  },
   validations () {
    return {
      userid: { required }, // Matches this.firstName
      name: { required }, // Matches this.lastName
      email: { required,email} // Matches this.contact.email
      }
    },
  // this for edit 
   created: function(){
            this.getItem();
    },


// Methods 
  methods:{
    // call the funtion 
    refreshData(){
        this.api.get("users/")
        .then((response)=>{
            this.user=response.data;
            // this.departmentsWithoutFilter=response.data;
        });
    },

//  Send the create record 
    Post_data_for_user(){
      
      this.api.post("/createuser",this.posts)
      .then((result)=>{
        this.result = result.data
        alert(result.data);
      })

     ///// Here catch the error
      .catch(error =>console.log(error))
      this.v$.$validations
      if (!this.v$.$error){
          alert('Form successfully submitted')
      }else{
          alert('Form failed validation')
      }

      // let up = this.posts
      // up.append('image',this.imageData,this.imageData.name) //gs://systolic-web-template.appspot.com/Profile
      // this.api.post("gs://systolic-web-template.appspot.com/Profile",up)
      //  .then((result)=>{
      //   this.result = result.data
      //   alert(result.data);
      // })

      //   .catch(error =>console.log(error))
      // this.v$.$validations
      // if (!this.v$.$error){
      //     alert('Form successfully submitted')
      // }else{
      //     alert('Form failed validation')
      // }
      // e.preventDefault();
      // alert(result.data);
      console.log("home")
    },


    back:function(){
      //  console.log(getlocal,'IDD')
       this.$router.push('/system-security/users')
    
    },

    chooseImage () {
      this.$refs.fileInput.click()
    },

// Select and upload Profile Image

onSelectFile () {
  const input = this.$refs.fileInput
  const files = input.files

  console.log(files,"yoooo")
  
  if (files && files[0]) {
    const reader = new FileReader

    reader.onload = e => {
    this.imageData = e.target.result
    // have to store in localstorage 
    let getlocal = this.imageData
    localStorage.setItem("image",getlocal)
     return getlocal
    }  

    reader.readAsDataURL(files[0])
    this.$emit('input', files[0])
  }
  else
  {
    alert('Sorry, FileReader API not supported')
  }
  

  
},

    
    updatedata(up){
      this.api.put("/user/"+this.id,{
        userid:this.userid,
        email:this.email,
        name:this.name,
      })
      .then((response)=>{alert(response.data)});
    
    },
    
  // for show entered Password string  
    show_password:function(){
          // const togglePassword = document.querySelector('#togglePassword');
          const password = document.querySelector('#password');
          const type = password.getAttribute('type') === 'password' ? 'text' : 'password';
          password.setAttribute('type', type);
            // toggle the eye / eye slash icon
          this.classList.toggle('bi-eye');
      },

      change({ coordinates, canvas }) {
			console.log(coordinates, canvas);
		  },
},
  
  mounted:function(){
    this.refreshData();
  },
 
}
</script>

<style scoped>
.profile{
  width:200px;
  height:200px;
  align-self: center;
  padding: 20px;

}
/* #file {
 display: none;
} */
.image{
  position: absolute;
  text-align: center;
  bottom: 25px;
  transform: translateX(-50% );
  height: 40px;
  width: 10%;
  left: 50%;
  line-height: 50px;
}



.fileup {
 display: none;
}

.profiler{
 height: 200px;
    width: 200px;
    /* position:absolute; */
    top: 30%;
    left: -33%;
    transform: translate(40%,-10%);
    border-radius: 50%;
    overflow: hidden;
    /* border: 1px solid grey; */
  
}
.preview-image{
  width: 200px;
  height: 200px;

}

.base-image-input {
  display: block;
  width: 200px;
  height: 200px;
  cursor: pointer;
  background-size: cover;
  background-position: center center;
}.placeholder {
  background: #F0F0F0;
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  color: #333;
  font-size: 18px;
  font-family: Helvetica;
}.placeholder:hover {
  background: #E0E0E0;
}.file-input {
  display: none;
}
</style>
