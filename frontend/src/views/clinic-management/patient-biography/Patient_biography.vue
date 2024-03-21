<template v-for="(color, key) in ['success', 'primary', 'info', 'warning', 'danger']">

  <div>
       <CToaster position="top-center">
      <CToast :show="alertboxSuccessIM" color="info" title="CToast static component">
             ProfileImage upload Successful!
      </CToast>
      <CToast :show="alertboxFailedIM" title="CToast static component">
        Failed Submit - {{ err }}
      </CToast>
    </CToaster>

     <CToaster position="top-center">
      <CToast :show="alertboxSuccess"  color="info" title="CToast static component">
        Document No. Were not select ! 
      </CToast>
      <!-- <CToast :show="alertboxFailed" title="CToast static component">
        Failed Submit - {{ err }}
      </CToast> -->
    </CToaster>
    <CRow> </CRow>
    <CRow>
      <CCol col="12">
        <CCard style="height: 100%">
          <!-- {{this.items}} -->
          <CCardBody>
            <!-- <div v-for="Info in PI" v-bind="Info.Id"> -->
            <CRow style="margin-top: -7px">
              <!-- <div v-for="data in Profile"> -->
                  <CCol col="4" align="left" style="display: flex">
                <label class="col-md-6 col-form-label"
                  ><strong> Registration No</strong></label
                >
                <!-- {{Search}} -->
                <CInput v-model="Search"  @input="PatientSearch()"   id="ShowRegistrationNo" />
                 <!-- <label for="ShowRegistrationNo" id="RegistrationNo">ppp</label> -->
              </CCol>
              <CCol col="5" align="left" >
                <div v-for="data in Profile" :key="data">
                    <label class=" col-form-label">
                       <div class="">
                  <p class="d-flex"><strong> {{data.PatientName}}</strong></p>
                      </div>
                  </label>
                   
                </div>
                
                <!-- <strong style="width: 12rem;"></strong> -->
              </CCol>

              <CCol  align="left" style="margin-left: -59px; ">
               <label class="col-md-6 col-form-label">
                 <div class="">
                 <p class="d-flex"><strong>DocumentNo</strong>&nbsp;:&nbsp;{{DOC_ID}}</p>
                 </div>
                </label>
                
                <!-- <strong style="width: 12rem;"></strong> -->
              </CCol>
              <!-- </div> -->
              
            </CRow>

            <CRow>
              <CCol
                align="center"
                class="p-1 ml-1 mr-1"
                style="margin-top: -20px"
              >
                <hr class="m-0 p-0" />
                <!-- <CCardHeader></CCardHeader> -->
                <CNav justified variant="tabs">
                  <CNavItem id='HOME' style="background-color:#f2f2f3;"
                    ><p class="m-0 p-0" v-on:click="Home">
                      <CIcon name="cil-home" size="lg" /> Home
                    </p></CNavItem
                  >
                  <CNavItem id='CLINIC_NOTE'
                    ><p class="m-0 p-0" v-on:click="Patient_Note">
                      <CIcon name="cil-file" size="lg" />Clinic Notes
                    </p></CNavItem
                  >
                  <CNavItem id='INVESTIGATION'
                    ><p class="m-0 p-0" v-on:click="Lab_investication">
                      <CIcon name="cil-people" size="lg" />Investigations
                    </p></CNavItem
                  >
                  <CNavItem id='DRUG'
                    ><p class="m-0 p-0" v-on:click="Radiology">
                      <CIcon name="cil-drop" size="lg" />Drugs Admin
                    </p></CNavItem
                  >
                  <CNavItem id='DIAGNOSIS'
                    ><p class="m-0 p-0" v-on:click="clinic_admissions">
                      <CIcon name="cil-layers" size="lg" />Diagnosis
                    </p></CNavItem
                  >
                </CNav>
              </CCol>
            </CRow>
            <div v-if="taps === 'default'">
              <CRow align="center" style="margin-left: -19px; margin-top: -4px">
                <CCol col="9">
                  <CDataTable
                    header
                    hover
                    striped
                    :items="visit"
                    :fields="visit1"
                    :items-per-page="11"
                    clickable-rows
                    :active-page="activePage"
                    @row-clicked="rowClickedForEncouter"
                    :pagination="{ doubleArrows: false, align: 'center' }"
                    @page-change="pageChange"
                    scopedSlots
                    sorter
                    @loadstart="load"
                  >
                

                  </CDataTable>
                </CCol>

                <CCol
                  align="center"
                  style="margin-left: -17px"
                  class="pl-1 pr-2"
                >
                  <CCard class="border-info shadow">
                    
                   <img  v-if="Photo !== null"
                      v-bind:src="Photo"
                      v-on:click="openUploadProfile"
                      class="
                        card-img-top
                        mx-auto
                        shadow
                        p-1
                        mb-1
                        bg-body
                        border-info
                        rounded-circle
                      "
                      style="width: 150px; height: 150px"
                      alt="..."
                    />
                    <img  v-else
                      :src="imagePreview"
                      v-on:click="openUploadProfile"
                      class="
                        card-img-top
                        mx-auto
                        shadow
                        p-1
                        mb-1
                        bg-body
                        border-info
                        rounded-circle
                      "
                      style="width: 150px; height: 150px"
                      alt="..."
                    />
                     <input
                      type="file"
                      accept="image/*"
                      name="image"
                      class="bg-info"
                      id="file"
                      v-on:change="UploadPreview"
                    style="display: none"
                  />
                    <div v-for="data in Profile" :key="data">
                    <div class="card-body m-0">
                      <div>
                        <CIcon
                          name="cil-location-pin"
                          class="text-primary"
                          width="13%"
                        />
                      </div>
                      <CRow
                        class="
                          card-text
                          text-center
                          flex
                          border-info
                          shadow
                          rounded
                        "
                        style="
                          margin-left: -1px;
                          margin-right: 1px;
                          margin-top: 10px;
                        "
                      >
                        <CCol>
                          <em>
                            <div style="font-size: 13px">
                               {{data.Address}}
                            </div></em
                          >
                        </CCol>
                      </CRow>

                      <CRow
                        class="shadow border-info rounded"
                        style="
                          margin-left: -1px;
                          margin-right: 1px;
                          margin-top: 2px;
                          height: 30px;
                        "
                      >
                        <CCol style="margin-right: -60px; margin-top: 3px"
                          ><CLable align="left" style="margin-left: 5px"
                            >Phone No</CLable
                          ></CCol
                        >
                        <CCol style="margin-top: 3px"
                          ><CLable style="margin-right: 9px"
                            >{{data.vcTelephone}}</CLable
                          ></CCol
                        >
                      </CRow>
                      <CRow
                        class="shadow border-info rounded"
                        style="
                          margin-left: -1px;
                          margin-right: 1px;
                          margin-top: 2px;
                          height: 30px;
                        "
                      >
                        <CCol style="margin-right: -50px; margin-top: 3px"
                          ><CLable align="left" style="margin-right: -42px"
                            >DOB</CLable
                          ></CCol
                        >
                        <CCol style="margin-top: 3px"
                          ><CLable style="margin-right: 19px"
                            >{{ new Date(data.dtpDateOfBirth).toLocaleDateString()}}</CLable
                          ></CCol
                        >
                      </CRow>
                      <CRow
                        class="shadow border-info rounded"
                        style="
                          margin-left: -1px;
                          margin-right: 1px;
                          margin-top: 2px;
                          height: 30px;
                        "
                      >
                        <CCol style="margin-right: -50px; margin-top: 3px"
                          ><CLable style="margin-right: -46px"
                            >NIC</CLable
                          ></CCol
                        >
                        <CCol style="margin-top: 3px"
                          ><CLable style="margin-right: 16px; margin-top: 3px"
                            >{{data.vcNICNo}}</CLable
                          ></CCol
                        >
                      </CRow>
                      <CRow
                        class="shadow border-info rounded"
                        style="
                          margin-left: -1px;
                          margin-right: 1px;
                          margin-top: 2px;
                          height: 30px;
                        "
                      >
                        <CCol style="margin-right: -50px; margin-top: 3px"
                          ><CLable style="margin-right: -46px">
                            Age</CLable
                          ></CCol
                        >
                        <CCol style="margin-top: 3px"
                          ><CLable style="margin-right: 76px">{{data.Age}}</CLable></CCol
                        >
                      </CRow>
                      <CRow
                        class="shadow border-info rounded"
                        style="
                          margin-left: -1px;
                          margin-right: 1px;
                          margin-top: 2px;
                          height: 30px;
                        "
                      >
                        <CCol style="margin-right: -50px; margin-top: 3px"
                          ><CLable style="margin-right: -25px"
                            >Gender</CLable
                          ></CCol
                        >
                        <CCol style="margin-top: 3px"
                          ><CLable style="margin-right: 49px"
                            >{{data.Gender}}</CLable
                          ></CCol
                        >
                      </CRow>
                    </div>
                    </div>
                  </CCard>
                </CCol>
              </CRow>
            </div>

            <div v-else-if="taps === 'laboratory'">
              <CCol>
                <!-- laboratory -->
                <CDataTable
                  header
                  hover
                  striped
                  :items="admis"
                  :fields="admission"
                  :items-per-page="10"
                  clickable-rows
                  :active-page="activePage"
                  @row-clicked="rowClicked"
                  :pagination="{ doubleArrows: false, align: 'center' }"
                  @page-change="pageChange"
                  column-filter
                  table-filter
                  sorter
                  items-per-page-select
                  @loadstart="load"
                >
                </CDataTable>
              </CCol>
            </div>

            <div v-else-if="taps === 'drug_admin'">
              heelo
              
            </div>

            <div v-else-if="taps === 'Radiology'">
              <CCol>
                <CDataTable
                  header
                  hover
                  striped
                  :items="radio"
                  :fields="radiology"
                  :items-per-page="10"
                  clickable-rows
                  :active-page="activePage"
                  @row-clicked="rowClicked"
                  :pagination="{ doubleArrows: false, align: 'center' }"
                  @page-change="pageChange"
                  column-filter
                  table-filter
                  sorter
                  items-per-page-select
                  @loadstart="load"
                >
                </CDataTable>
              </CCol>
            </div>

            <div v-else-if="taps === 'Admissions'">
              <CCol>
                <!-- Diagnozise -->
              <CDataTable
                              style="
                                text-align: left;
                                margin-top: -5px;
                                height: 100%;
                              "
                              header
                              clickableRows
                              :tableProps="{
                                striped: true,
                                hover: true,
                              }"
                              hover
                              striped
                              :items="GetDiagnisisinfo"
                              :fields="fields"
                              :items-per-page="10"
                              @row-clicked="show"
                              :pagination="{
                                doubleArrows: false,
                                align: 'center',
                              }"
                            >
                             <!-- <div> -->
                              <template #show_details="{ item }" class="">
                                <!-- d{{item.id}}  | clickable-rows -->
                                <td class="py-10">
                                  <CButtonGroup align="right">
                                    <div
                                      style=""
                                      size="sm"
                                      class="mt-6 d-flex"
                                      @click="show(item)"
                                    >
                                      <i @click="edit(item)"
                                        ><CIcon
                                          name="cil-pencil"
                                          size="sm"
                                          class="ml-2 mb-2 text-primary"
                                      /></i>

                                      <i @click="deleteinfo(item)"
                                        ><CIcon
                                          name="cil-trash"
                                          size="sm"
                                          class="ml-2 mb-2 text-danger"
                                      /></i>
                                    </div>
                                  </CButtonGroup>
                                </td>
                               
                                  
                                
                              </template>
                              <template #ProblemDescription_list="{ item }" class="">
                                <!-- d{{item.id}}  | clickable-rows -->
                                <td class="py-10">
                                  <CButtonGroup align="right">
                                    <div
                                      style="mar"
                                      size="sm"
                                     
                                      @click="show(item)"
                                    >
                                      <p class="text-left">{{item.ProblemDescription}}</p>
                                      <!-- <span class="badge badge-success">Success</span> {{item.vcSpeciality}} -->
                                      <p class="text-left" style="margin-top:-15px; font-size: 12px;">
                                       
                                        <b class="badge badge-secondary bg-info " style="color:black;">Medical Clinic</b></p>
                                    </div>
                                  </CButtonGroup>
                                </td>
                              </template>
                              <!-- </div> -->
                            </CDataTable>
              </CCol>
            </div>

            <!-- </div> -->
          </CCardBody>

          <!-- <div v-if="taps === 'default' ">
          <CCard md="12">
            LAB Investication
          </CCard>
        </div> -->
        </CCard>
      </CCol>
    </CRow>
    <!-- </CContainer> -->
  </div>
</template>

<script>
import Laboratory from "@/views/clinic-management/lab_investigation/Laboratory_investigation";
import { DateTime } from 'luxon'
import * as Vue from "vue";
import axios from "axios";
import VueAxios from "vue-axios";

export default {
  // cilPencil
  name: "Users",
  inject: ["api"],
  // props:[items],
  data() {
    return {
      err:'',
      alertboxSuccess:false,
      Profile:null,
      Search:localStorage.getItem("RegiNo"),
      // Search:null,
      items: [],
      GetDiagnisisinfo:null,
      alertboxSuccessIM:false,
      alertboxFailedIM:false,
      Photo:null,
      index: null,
      DOC_ID:null,
      imagePreview: "/img/avatars/blank.png",
      PI: [],
      dataPA: [],
      visit: [],
      fields: [
        {
          filter: false,
          sorter: false,
          key: "ProblemDescription_list",
          label: "Problem",
          _style: { width: "14%" },
        },
        {
          filter: false,
          sorter: false,
          key: "Note",
          label: "Note",
          _style: { width: "65%" },
        },
        {
          filter: false,
          sorter: false,
          key: "Status",
          label: "Status",
          _style: { width: "10%" },
        },
        
      ],
      radio: [
        {
          id: 2,
          date: "2021-11-27",
          test_code: "A0081",
          description: "Condition is good",
        },
        {
          id: 3,
          date: "2021-10-27",
          test_code: "DC008110254",
          description: "Have check urine track drink more more Hydrine",
        },
      ],
      admis: [
        {
          id: 2,
          date: "2020-10-02",
          Admintion: "High Fever back pain",
          Admintion_No: "20252/0251/698",
        },
        {
          id: 3,
          date: "2021-07-14",
          Admintion: "Unfitness and high cough",
          Admintion_No: "20252/0251/698/365",
        },
      ],
      taps: "default",
      radiology: [
        { filter: false, sorter: false, key: "date", label: " Date" },
        { filter: false, sorter: false, key: "test_code", label: "Test code" },
        {
          filter: false,
          sorter: false,
          key: "description",
          label: "Descriptions",
        },
      ],
      admission: [
        { filter: false, sorter: false, key: "date", label: "Date" },
        { filter: false, sorter: false, key: "Admintion", label: "Admittance" },
        {
          filter: false,
          sorter: false,
          key: "Admintion_No",
          label: "Admittance No",
        },
      ],
      visit1: [
        { filter: false, sorter: false, key: "index", label: "#" },
        {
          filter: false,
          sorter: false,
          key: "encouterId",
          label: "Encounter ID",
        },
        {
          filter: false,
          sorter: false,
          key: "Serlocation",
          label: "Service Location",
        },
        { filter: false, sorter: false, key: "date", label: "Date" },
      ],
    };
  },
  computed: {
    index() {
      let se = this.visit.length;
      return (index = se);
    },
  },

  watch: {
    $route: {
      immediate: true,
      handler(route) {
        if (route.query && route.query.page) {
          this.activePage = Number(route.query.page);
        }
      },
    },
  },

  methods: {
    back: function () {
      this.$router.push("/Clinic Management/doctor-queue-or-portal");
    },

  /*================================================================================*/ 
  //                         UPLOAD THE PROFILE
  /*================================================================================*/ 
    openUploadProfile(){
       document.getElementById("file").click();
    },
    /**/
    UploadPreview(e){
      var reader,
      files = e.target.files;
      reader = new FileReader();
      reader.onload = (e) => {
        let img = e.target.result;

        this.Photo = img;
    }
    reader.readAsDataURL(files[0]);
   
    setTimeout((this.UploadProfileImage),1000)

    },
 /*=================================================================================*/
 //                          PATIENT ENCOUNTER INFO LIST
 /*=================================================================================*/
    async getAppointments() {
      let regcode = this.$route.params.id;
      const result = await this.api.get(
        `doctorportal/patientactivityinfo?regcode=${regcode}`
      );
      let patient,
        visit = result.data.PatientActivityInfo;
      patient = visit.map((value, data) => ({
        index: data + 1,
        date: new Date(value.dtpCreateDate).toLocaleDateString("en-GB"),
        encouterId: value.cBHTNo,
        Serlocation: value.cDiagnosisCode,
        DocumentNo:value.DocumentNo
      }));
      // this.visit = patient;
       var uo = patient.sort((a, b) => DateTime.fromFormat(b.date,'dd/MM/yyyy') - DateTime.fromFormat(a.date,'dd/MM/yyyy') )
       this.visit = uo.map((value,data)=>({index:data+1,date:value.date,encouterId:value.encouterId,Serlocation:value.Serlocation,DocumentNo:value.DocumentNo}))
      this.GetAutoloadFORdocID()
      this.GetDiagnosisdata();
    },
  

    /*======================================================================================*/
    //                          PATIENT PROFILE DATA
    /*======================================================================================*/
    async Promise(){
      const RegiNumber = this.$route.params.id;
      localStorage.setItem("RegiNo", RegiNumber.trim());
      let RegiNo = localStorage.getItem("RegiNo")
      const resultGetPatientRegistionInfo = await this.api.get(`/reception/GetPatientRegistrationInfo?RegiNo=${RegiNo}`);
      let data = resultGetPatientRegistionInfo.data.GetPatientRegistrationInfo
       if (data !== null) {
        switch (data[0].Gender) {
          case "M":
            data[0].Gender = "Male";
            break;
          case "F":
            data[0].Gender = "Female";
            break;

          default:
            break;
        }
      this.Profile = data
    }
    },

    /*======================================================================================*/
    //                         PATIENT PROFILE IMAGE
    /*======================================================================================*/
    async GetProfile(){
      let RegiNo = localStorage.getItem("RegiNo")
      const result = await this.api.get(`Clinic-notes/GetPatientProfileImage?RegiNo=${RegiNo}`);
      let ProfileImage = result.data.GetPatientProfileImage
      this.Photo =ProfileImage[0].Patient_Image
       
    },
    
    /*=====================================================================================*/
    //                             LATEST DOC NO LOAD                                      //
    /*=====================================================================================*/
      async GetAutoloadFORdocID(){
      let regi, clinic;
      regi = String(localStorage.getItem("RegiNo"));
      clinic = String(localStorage.getItem("Location"));
      const result = await this.api.get( `/Clinic-notes/LastWiseClinicSessions?regno=${regi}&clinic=${clinic}`
      );
      let Data,LastDOCNum = result.data.RegistrationWiseClinicSessions;
      if(LastDOCNum != null){
          Data = LastDOCNum[0].DocumentID.trim()
          localStorage.setItem("DDOCUMENT",Data)
          this.DOC_ID = localStorage.getItem("DDOCUMENT")
          this.GetDisplayDocumentNo();
      } else{
          localStorage.setItem("DDOCUMENT","undefined")
      }
         
      },

    /*======================================================================================*/
    //                        Click EncouterRow Get Document No.
    /*======================================================================================*/

    rowClickedForEncouter(item){
      
      if(item.DocumentNo === "null" || item.DocumentNo === "undefined" || item.DocumentNo === ""){
        localStorage.setItem("DDOCUMENT","null")
        // this.GetDisplayDocumentNo();
      } 
      else{
        localStorage.setItem("DDOCUMENT",item.DocumentNo)
        this.GetDisplayDocumentNo();
      }
     
    },


    /*======================================================================================*/
    //                        URl redirect to Clinic Note
    /*======================================================================================*/
    Patient_Note: function () {
      let DocumentNo = localStorage.getItem("DDOCUMENT")
      if( DocumentNo === "null" || DocumentNo==='undefined'){
        this.alertboxSuccess = true
        setTimeout(this.fadeoff, 3000);
        
      }
      else{
          let route_params = this.$route.params.id;
          let route = this.$router.resolve({path: `/doctor-queue-or-portal/${route_params}/Notes`,});
          window.open(route.href, '_blank');
      }
     
    },

    fadeoff:function(){this.alertboxSuccess = false},

     /*======================================================================================*/
    //     NAVIGATION TABS | HOME | CLINIC-NOTES | INVESTIGATIONS | DRUGS-ADMIN | DIAGNOSIS | 
    /*======================================================================================*/
    Lab_investication: function () {
      let route_params = this.$route.params.id;
      this.taps = "laboratory";
    },
    Drugs_admin: function (item) {
      this.taps = "drug_admin";
    },
    Radiology: function (item) {
      this.taps = "Radiology";
    },
    clinic_admissions: function (item) {
      this.taps = "Admissions";
    },
    // Patient biography
    Home: function (item) {
      this.taps = "default";
      document.getElementById("HOME").style.color = "blue";

    },


    /*======================================================================================*/
    //                      REGISTRATION NUMBER WISE PATIENT SEARCH
    /*======================================================================================*/
    async PatientSearch(){
      let regcode = this.Search
      const result = await this.api.get(`doctorportal/patientactivityinfo?regcode=${regcode}`);
      let patient,
        visit = result.data.PatientActivityInfo;
        patient = visit.map((value, data) => ({
        index: data + 1,
        date: new Date(value.dtpCreateDate).toLocaleDateString("en-GB"),
        encouterId: value.cBHTNo,
        Serlocation: value.cDiagnosisCode,
        DocumentNo:value.DocumentNo
      }));
      
      setTimeout( () => this.$router.push({path:`/doctor-queue-or-portal/${regcode}`}),1000);
  
    },

    /*======================================================================================*/
    //                    AUTOUPLOAD PATIENT PROFILE IMAGE
    /*======================================================================================*/
    async UploadProfileImage(){
      let form ={}
      form.RegiNo = String(localStorage.getItem("RegiNo"))
      form.Patient_Image = this.Photo
      form.User_ID = String(localStorage.getItem("UserId"))
      this.api
        .post("Clinic-notes/SaveProfileImage", form)
        .then((result) => {
          if (result.data === "success") {
            this.alertboxSuccessIM = true;
            setTimeout(() => (this.alertboxSuccessIM = false), 3000);
            this.GetProfile();
          } else {
            this.alertboxFailedIM = true;
            setTimeout(() => (this.alertboxFailedIM = false), 3000);
          }
        })
        .catch((error) => {
          this.err = error;
          this.alertboxFailed = true;
          setTimeout(() => (this.alertboxFailedIM = false), 3000);
        });
    },
    /*==================================================================*/
    //           SHOW THE PATIENT DOCUMENT NUMBER                      //
    /*==================================================================*/
   async GetDisplayDocumentNo(){
      this.DOC_ID = localStorage.getItem("DDOCUMENT")
    },
    /*==================================================================*/
    //                 SHOW THE DIAGNOSIS                                     //
    /*==================================================================*/
    async GetDiagnosisdata(){
       let DocNo = localStorage.getItem("DDOCUMENT");
       const result = await this.api.get(`/Clinic-notes/GetDocumentHeaderProblemList?DocID=${DocNo}`)
       this.GetDiagnisisinfo = result.data.GetDocumentHeaderProblemList
    }

  },
  created() {
    this.getAppointments();
    this.Promise();
    this.GetProfile();
    this.GetDocHeaderProblemlists();
    this.ShowRegistration();
    this.GetDisplayDocumentNo();
  },
};
</script>
<style>
</style>
