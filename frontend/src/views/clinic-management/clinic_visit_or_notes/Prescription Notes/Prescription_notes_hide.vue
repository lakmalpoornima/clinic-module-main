<template>
  <div>

     <!-- This is for referral Template Name  -->
      <CModal
      title="Referral Template"
      color="warning"
      :show.sync="ReferraltemplateName" class="align-self-center"
    >
     <CRow>
       <CCol>
         <CInput type="text" label="Template Name" v-model="ReferralTemp"/>
       </CCol>
     </CRow>
      <div slot="footer">
        <CRow >
       <CCol  class="p-1">
       <CButton class="col align-self-center" color="primary" @click="SaveasTemplate">Save</CButton>          
       </CCol>
       <CCol  class="p-1">
       <CButton class="col align-self-center" color="secondary" @click="deleteDiagnosisModal = false">Cancel</CButton>
       </CCol>
        </CRow>
      </div>
    </CModal>
    <!--  -->

    <CToaster position="top-center">
      <CToast :show="alertboxSuccess" title="CToast static component">
        Successful Submit
      </CToast>
      <CToast :show="alertboxFailed" title="CToast static component">
        Failed Submit - {{ err }}
      </CToast>
    </CToaster>
     <CModal
      title="Delete Diagnosis"
      color="warning"
      :show.sync="deleteDiagnosisModal" class="align-self-center"
    >
      Confirm to delete Dignosis item ?
      <div slot="footer">
        <CRow >
       <CCol  class="p-1">
       <CButton class="col align-self-center" color="primary" @click="deleteSpecialNote(deleteDiagnosisItem,problemList)">Confirm</CButton>          
       </CCol>
       <CCol  class="p-1">
       <CButton class="col align-self-center" color="secondary" @click="deleteDiagnosisModal = false">Cancel</CButton>
       </CCol>
        </CRow>
      </div>
    </CModal>
    <CModal
      title="Delete Prescription"
      color="warning"
      :show.sync="deletePrescriptionModal" class="align-self-center"
    >
      Confirm to delete Prescription item ?
      <div slot="footer">
        <CRow >
       <CCol  class="p-1">
       <CButton class="col align-self-center" color="primary" @click="deletePrescription(deletePrescriptionItem)">Confirm</CButton>          
       </CCol>
       <CCol  class="p-1">
       <CButton class="col align-self-center" color="secondary" @click="deletePrescriptionModal = false">Cancel</CButton>
       </CCol>
        </CRow>
      </div>
    </CModal>
    <div>
      <template>
        <div>
          <CModal
            title="Load Template"
            color="warning"
            :show.sync="loadtemplate"
            size="lg"
          >
            <CDataTable
              header
              hover
              striped
              :items="ReferralLetterTemplate"
              :fields="ReferralLetterTemplateField"
              :items-per-page="10"
              clickable-rows
              :active-page="activePage"
              @row-clicked="rowClicked"
              :pagination="{
                doubleArrows: false,
                align: 'center',
              }"
              @page-change="pageChange"
              scopedSlots
              sorter
              @loadstart="load"
            >
              <template #call="{ item }">
                <td class="py-10">
                  <CButtonGroup>
                    <div size="sm" class="mt-6" @click="loadRefTemplate(item)">
                      <CButton
                        type="button"
                        color="behance"
                        size="sm"
                        class="font-weight-bold1"
                        >Load</CButton
                      >
                    </div>
                  </CButtonGroup>
                </td>
              </template>
            </CDataTable>

            <template #footer>
              <CButton @click="close" color="facebook">close</CButton>
            </template>
          </CModal>
          <!-- Referral all the Tamplate  loading END Here-->

          <!-- Comments image view start Here-->
          <CModal
            title="Patient disease history of Image (PDHI)"
            color="warning"
            :no-close-on-backdrop="true"
            :borderColor="yahoo"
            :show.sync="warningModal"
          >
            <div class="card" style="max-width: 800px">
              <div class="row g-0">
                <div class="col-sm-5">
                  <div v-for="data in ImageOfDisease" :key="data">
                    <img
                      v-bind:src="data.cNameofImage"
                      class="card-img-top rounded h-100"
                      alt="..."
                    />
                  </div>
                </div>
                <div class="col-sm-7">
                  <div class="card-body" style="margin-left: -25px">
                    <CLable color="info" class="card-title font-bolder"
                      ><b>Left Thumb Nail virus attack</b></CLable
                    >
                    <p class="card-text">
                      Nail fungal infections are the most common diseases of the
                      nails, making up about 50 percent of nail abnormalities.
                      Fungus is normally present on the body, but if it
                      overgrows, it can become a problem.
                    </p>
                    <div class="d-flex">
                      <p class="font-italic text-info text-wrap" color="info">
                        Examine by : Dr.Dhayananda
                      </p>
                      <p
                        class="font-italic text-yahoo text-wrap"
                        style="
                          font-size: 12px;
                          margin-top: 3px;
                          margin-left: 5px;
                        "
                      >
                        2021-03-14
                      </p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <template #footer style="margin-bottom: -300px" class="bg-primary">
              <CButton @click="close" color="facebook">close</CButton>
            </template>
          </CModal>
          <!-- Comments image view END Here-->
          <!--=======================================================================================-->
          <!-- Comments imageupload view start Here-->
          <CModal
            :show.sync="imageModal"
            title="Patient disease of sample upload"
            color="warning"
            bgcolor="danger"
            style="background-color: #3c4b64"
          >
            <CCard class="border border-info">
              <CRow>
                <CCol
                  style="
                    margin-left: 8px;
                    margin-top: 10px;
                    margin-bottom: 10px;
                    margin-right: 7px;
                  "
                >
                  <!--  -->
                  <img v-if="imagePreviewFordisease!== null"
                    id="output"
                    v-bind:src="imagePreviewFordisease"
                    class="border border-danger rounded"
                    v-on:click="openUpload"
                    style="width: 150px; height: 150px"
                  
                   />
                  <img v-else
                    id="output"
                    src="img/avatars/upload.png"
                    class="border border-danger rounded"
                    v-on:click="openUpload"
                    
                    style="width: 150px; height: 150px"
                  />
                  <!--  -->
                  <input
                    type="file"
                    accept="image/*"
                    name="image"
                    class="bg-info"
                    id="file"
                    v-on:change="UploadPreview"
                    style="display: none"
                  />
                </CCol>
                <CCol
                  style="
                    margin-left: -180px;
                    margin-right: 2px;
                    margin-bottom: 2px;
                  "
                >
                  <CRow style="margin-top: 4px">
                    <CCol class="" align="left">
                      <CLable>Disease of name</CLable> <CInput type="text" />
                    </CCol>
                  </CRow>
                  <CRow style="margin-top: -19px">
                    <CCol class="" align="left">
                      <CLable>Note </CLable> <CTextarea type="text" />
                    </CCol>
                  </CRow>
                  <CRow style="margin-top: -15px">
                    <CCol
                      class="d-flex"
                      align="center"
                      style="margin-left: 10px"
                    >
                      <CButton
                        @click="close"
                        style="width: 8rem; margin-left: 10px"
                        color="facebook"
                        >clear</CButton
                      >
                      <CLoadingButton
                        progress
                        :spinner="false"
                        timeout="2000"
                        @click="close"
                        style="width: 8rem; margin-left: 10px"
                        color="yahoo"
                        >Save</CLoadingButton
                      >
                    </CCol>
                  </CRow>
                </CCol>
              </CRow>
            </CCard>

            <template #footer>
              <!-- <CButton @click="clickok" color="danger">Discard</CButton> -->
              <CButton @click="closeIm" color="facebook">close</CButton>
            </template>
          </CModal>
          <!-- Comments imageupload view END Here-->
        </div>
        <!-- All the modal are ENDING here the link -->
      </template>

      <!--{ All the content is START HERE}-->
      <CRow>
        <CCol md="3">
          <CCard class="health-check-table" style="height: 100%">
            <!-- Profile Start -->
            <CCol class="pl-1 pr-2">
              <CCard
                style="
                  margin-left: -3px;
                  margin-right: -7px;
                  margin-bottom: 9px;
                "
              >
                <!-- <CButton >gt</CButton> -->
                <!--  -->
                <div v-for="data in Profile" :key="data">
                  <div class="d-flex flex-row mb-3" key="data.id">
                    <img v-if="Photo !== null"
                      class="text-center m-0 shadow rounded"
                      v-bind:src="Photo"
                      width="85"
                      alt=""
                      
                    />
                    <img v-else
                      class="text-center m-0 shadow rounded"
                      :src="imagePreview"
                      width="75"
                      :alt="imagePreview"
                      style="height:90px;"
                    />
                    <div class="d-flex flex-column ml-1">
                      <CRow
                        style="margin-top: 2px; margin-left: -4px"
                        class="rounded"
                      >
                        <CCol
                          ><CLable
                            >Name<span
                              style="margin-left: 13.5px; margin-right: 4px"
                            >
                              :</span
                            >
                            <span class="text-primary">{{ data.PatientName }}</span>
                          </CLable></CCol
                        >
                      </CRow>
                       <CRow
                        style="margin-top: 1px; margin-left: -4px"
                        class="rounded"
                      >
                        <CCol
                          ><CLable
                            >Registration No<span class="ml-2" style="margin-right: 4px">
                              :</span
                            ><span class="text-primary">{{
                              data.RegistrationNo
                            }}</span></CLable
                          ></CCol
                        >
                      </CRow>
                      <CRow
                        style="margin-top: 2px; margin-left: -4px"
                        class="rounded"
                      >
                        <CCol
                          ><CLable
                            >NIC<span
                              style="margin-left: 13.5px; margin-right: 4px"
                            >
                              :</span
                            >
                            <span class="text-primary">{{ data.vcNICNo }}</span>
                          </CLable></CCol
                        >
                      </CRow>
                     <CRow
                        style="margin-top: 1px; margin-left: -4px"
                        class="rounded"
                      >
                        <CCol
                          ><CLable
                            >DOB<span class="ml-2" style="margin-right: 4px">
                              :</span
                            ><span class="text-primary">{{
                              new Date(data.dtpDateOfBirth).toLocaleDateString()
                            }}</span></CLable
                          ></CCol
                        >
                      </CRow>
                      <CRow
                        style="margin-top: 1px; margin-left: -4px"
                        class="rounded"
                      >
                        <CCol
                          ><CLable
                            >Age<span class="ml-3" style="margin-right: 5px"
                              >:</span
                            ><span class="text-primary">{{
                              data.Age
                            }}</span></CLable
                          ></CCol
                        >
                      </CRow>
                    </div>
                  </div>
                  <CRow style="margin-top: -16px" class="">
                    <CCol
                      class="p-0 text-left rounded"
                      style="margin-left: 14px; margin-right: -40px"
                      ><CLable>
                        &nbsp;&nbsp;&nbsp;&nbsp;Gender<span
                          style="margin-left: 14px; margin-right: 5px"
                          >:</span
                        ><span class="text-primary">Female</span></CLable
                      ></CCol
                    >
                  </CRow>
                  <CRow style="margin-top: -1px" class="">
                    <CCol
                      class="p-0 text-left rounded"
                      style="margin-left: 14px; margin-right: -40px"
                      ><CLable>
                        &nbsp;&nbsp;&nbsp;&nbsp;Tel No<span
                          style="margin-left: 20px; margin-right: 4px"
                          >:</span
                        ><span class="text-primary">{{
                          data.vcTelephone
                        }}</span></CLable
                      ></CCol
                    >
                  </CRow>
                  <CRow style="margin-top: -1px" class="d-flex">
                    <CCol
                      ><CLable style="margin-left: 14px; margin-right: -40px"
                        >Address<span
                          style="margin-left: 10px; margin-right: 4px"
                          >:</span
                        >
                        <address
                          class="text-left text-primary font-italic"
                          style="
                            margin-top: -15px;
                            margin-left: 80px;
                            width: 54%;
                            height: 54px;
                          "
                        >
                          {{ data.Address }}
                        </address></CLable
                      ></CCol
                    >
                  </CRow>
                </div>
              </CCard>
              <!-- Profile End -->

              <!-- Date start -->
              <CCard
                style="
                  margin-bottom: -3px;
                  margin-top: -8px;
                  margin-left: -3px;
                  margin-right: -7px;
                  height: 6%;
                "
              >
                <CRow style="margin-bottom: -3px; margin-top: 3px">
                  <CSelect
                    class="col align-self-center"
                    :value.sync="ClinicWiseSession"
                    :options="ClinicSession"
                    disabled
                  ></CSelect>
                </CRow>
              </CCard>
              <!-- Date End -->

              <!-- Body info input-->
              <div
                class=""
                style="margin-top: -18px; margin-bottom: 1px; padding: 0px"
              >
                <CRow style="margin-top: 20px">
                  <CCol class="d-flex" align="center">
                    <!-- <CButton @click="show(item)">HU</CButton> -->
                    <CLable
                      class=""
                      color="dark"
                      style="margin-top: 10px; margin-left: 2px"
                      >BP</CLable
                    >
                    <!-- DE{{patient.RegistrationNo}} -->
                    <!-- {{DocumentHeaderNotes}} -->
                    <!-- <CInput style="width: 60%; margin-right: -40px; margin-left: 74px" v-model="DocumentHeaderNotes[0].BP"/> -->
                    <CInput
                      style="width: 60%; margin-right: -40px; margin-left: 74px"
                      v-model="header.bp"
                    />
                  </CCol>
                </CRow>

                <CRow style="margin-top: -18px">
                  <CCol class="d-flex">
                    <CLable style="margin-top: 10px; margin-left: 2px"
                      >Pulse Rate</CLable
                    >
                    <CInput
                      type="number"
                      placeholder=""
                      class="mt-1"
                      style="
                        width: 60%;
                        margin-right: -31px;
                        margin-left: 26.4px;
                      "
                      v-model="header.pulse_rate"
                    />
                  </CCol>
                </CRow>

                <CRow style="margin-top: -18px" align="center">
                  <CCol class="d-flex">
                    <CLable style="margin-top: 10px; margin-left: 2px"
                      >Height</CLable
                    >

                    <CInput
                      type="number"
                      class="mt-1"
                      style="
                        width: 60%;
                        margin-right: -70px;
                        margin-left: 48.5px;
                      "
                      min="80"
                      max="150"
                      @input="calcBMI"
                      v-model="header.height"
                      valid-feedback="Input is valid."
                      invalid-feedback="Please Enter The Referral ID."
                    />
                  </CCol>
                </CRow>
                <CRow style="margin-top: -18px" align="center">
                  <CCol class="d-flex">
                    <CLable style="margin-top: 10px; margin-left: 2px"
                      >Weight(kg)</CLable
                    >
                    <!-- <CLable></CLable> -->
                    <!-- <CIcon name="cil-speedometer" class='mt-3 ml-1 mr-2 text-info' size='md'/> -->
                    <CInput
                      type="number"
                      placeholder=""
                      class="mt-1"
                      style="
                        width: 60%;
                        margin-right: -20px;
                        margin-left: 22.5px;
                      "
                      min="80"
                      max="150"
                      @input="calcBMI"
                      v-model="header.weight"
                      valid-feedback="Input is valid."
                      invalid-feedback="Please Enter The Referral ID."
                    />
                  </CCol>
                </CRow>

                <CRow style="margin-top: -18px" align="center">
                  <CCol class="d-flex">
                    <CLable style="margin-top: 10px; margin-left: 2px"
                      >BMI</CLable
                    >

                    <CInput
                      type="number"
                      class="mt-1"
                      style="width: 60%; margin-right: -70px; margin-left: 66px"
                      min="80"
                      max="150"
                      v-model="header.bmi"
                      valid-feedback="Input is valid."
                      invalid-feedback="Please Enter The Referral ID."
                    />
                  </CCol>
                </CRow>
                <CRow style="margin-top: -18px">
                  <CCol align="center" class="d-flex">
                    <CLable
                      style="margin-top: 10px; margin-left: 2px; padding: 0px"
                      >Classification</CLable
                    >
                    <CSelect
                      class="col align-self-center mt-1"
                      style="
                        width: 60%;
                        margin-right: -20.9px;
                        margin-left: -3px;
                      "
                      :value.sync="header.classificationValue"
                      :options="classificationDropdown"
                    >
                    </CSelect>
                    <!-- <CInput  label="Status" style="width:90%; margin-top:-5px" min="30" max="310" valid-feedback="Input is valid." invalid-feedback="Please Enter The Referral ID."/> -->
                  </CCol>
                </CRow>
                <CRow style="margin-top: -14px; margin-bottom: 3px">
                  <CCol align="center" class="d-flex">
                    <CLable
                      style="
                        margin-top: 10px;
                        margin-right: -4px;
                        margin-left: 2px;
                        padding: 0px;
                      "
                      >Status</CLable
                    >
                    <CSelect
                      class="col align-self-center mt-1"
                      style="width: 60%; margin-right: -21px; margin-left: 41px"
                      :value.sync="header.status"
                      :options="statusDropdown"
                    >
                    </CSelect>
                    <!-- <CInput  label="Status" style="width:90%; margin-top:-5px" min="30" max="310" valid-feedback="Input is valid." invalid-feedback="Please Enter The Referral ID."/> -->
                  </CCol>
                </CRow>

                <CRow style="margin-top: -12px; margin-bottom: 3px">
                  <CCol align="center">
                    <CInput
                      placeholder="Status Comment"
                      style="
                        width: 100%;
                        margin-top: -5px;
                        padding: 0;
                        margin-left: -10px;
                        margin-right: -14px;
                      "
                      min="80"
                      max="150"
                      valid-feedback="Input is valid."
                      v-model="header.status_com"
                      invalid-feedback="Please Enter The Referral ID."
                    />
                    <!-- {{BodyMI}} -->
                  </CCol>
                </CRow>
              </div>
            </CCol>
            <!-- </CCardBody>  -->
          </CCard>
        </CCol>

        <CCol md="9" class="health-check-table">
          <CCard style="margin-left: -28px; height: 100%">
            <CRow class="flex-row" style="">
              <CCol col="10" class="flex-row">
                <CTabs  variant="tabs" :activeTab.sync="tabIndex">
                  <CTab title="Diagnosis">

                    <CTabContent style="margin-top: 50px; padding: 0px margin-left:-30px;">
                      <!-- <CLable><b>Problem list</b></CLable> -->
                    <CTabPane >
                          <CRow style="padding: 0px;">
                            <CCol class="d-flex">
                              <CSelect  hidden
                                placeholder=" Select a ICD 10"
                                :value.sync="diagnosis.ProblemDescriptionValue"
                                :options="diagnosticList"
                              >
                              </CSelect>
                              <CInput hidden
                                v-model="diagnosis.ProblemNote"
                                placeholder="Note"
                                style="width: 70%"
                              />
                              <CInput hidden
                                type="number"
                                v-model="diagnosis.ProblemPriority"
                                min="0"
                                style="width: 20%"
                              />
                              <CInputCheckbox hidden
                                :checked.sync="diagnosis.ProblemStatus"
                              />
                                <CButton hidden
                                  size="md"
                                  color="info"
                                  @click="addSpecialNote(diagnosis,problemList)" disabled
                                  style="width: 8rem; height: 2rem"
                                  > {{ diagnosis.index === null ? 'Add' : 'Update'}}</CButton
                                >
                            </CCol>
                          </CRow>
                         
                              
                          
                          <CRow style="margin-top:-10px;">
                             <CCol>
                          <CDataTable
                            style=" padding: 5px;
                              text-align: left;
                              margin-top: -11px;
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
                            :items="problemList"
                            :fields="problemListFields"
                            :items-per-page="10"
                            :pagination="{
                              doubleArrows: false,
                              align: 'center',
                            }"
                          >
                            <!-- <div> -->
                            <template #show_details="{ item }" class="">
                              <!-- d{{item.id}}  | clickable-rows -->
                              <td class="py-10">
                                <CButtonGroup align="right" disabled>
                                  <div
                                    style=""
                                    size="sm"
                                    class="mt-6 d-flex"
                                  >
                                    <i  hidden style="cursor: no-drop;"
                                      ><CIcon
                                        name="cil-pencil"
                                        size="sm"
                                        class="ml-2 mb-2 text-primary"
                                    /></i>

                                    <i hidden style="cursor: no-drop;"
                                      ><CIcon
                                        name="cil-trash"
                                        size="sm"
                                        class="ml-2 mb-2 text-danger"
                                    /></i>
                                  </div>
                                </CButtonGroup>
                              </td>
                            </template>
                            <template
                              #ProblemDescription_list="{ item }"
                              class=""
                            >
                              <!-- d{{item.id}}  | clickable-rows -->
                              <td class="py-10">
                                <CButtonGroup align="right" disabled>
                                  <div
                                    style="mar"
                                    size="sm"
                                    @click="show(item)"
                                  >
                                    <p class="text-left">
                                      {{ item.ProblemDescription }}
                                    </p>
                                    <!-- <span class="badge badge-success">Success</span> {{item.vcSpeciality}} -->
                                    <!-- <p class="text-left" style="margin-top:-15px; font-size: 12px;">
                                       
                                        <b class="badge badge-secondary bg-info " style="color:black;">Medical Clinic</b></p> -->
                                  </div>
                                </CButtonGroup>
                              </td>
                            </template>
                            <!-- </div> -->
                          </CDataTable>
                           </CCol>
                          </CRow>
                         
                        </CTabPane>
                        </CTabContent>
                  </CTab>

                  <CTab title="Lab Investigation">
                    <CTabContent style="margin-top: 50px;padding: 0px;" align="center">
                     <CTabPane >
                          <CCol style="padding: 5px" class="d-flex" align="right">
                            <CButton hidden color="behance"> Load Data </CButton>
                            &nbsp;&nbsp;&nbsp;
                            <!-- <CButton color="reddit"> Clear </CButton> -->
                          </CCol>
                          <CDataTable
                            style="padding: 5px;
                              text-align: left;
                              margin-top: 8px;
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
                            :items="LabInvestData"
                            :fields="LabInvestfields"
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
                            <template #inputfield="{ item }" class="">
                              <!-- d{{item.id}}  | clickable-rows -->
                              <td class="py-10">
                                <CButtonGroup align="left">
                                  <div style="mar" size="sm">
                                    <CInput
                                      placeholder="Value"
                                      :checked="item._selected"
                                      @update:checked="() => check(item)"
                                      style="width: 40%"
                                    />
                                  </div>
                                </CButtonGroup>
                              </td>
                            </template>
                            <!-- </div> -->
                          </CDataTable>
                        </CTabPane>
                      </CTabContent>
                  </CTab>

                  <CTab title="Prescriptions">
                    <CTabContent style="margin-top: 50px; padding: 0px">
                    <CTabPane>
                          <CCardBody >
                            <CRow align="center" style="margin-top: -1px;">

                              <CCol
                                col="5"
                                style=""
                              >
                              <!-- <CButton @click="ad">TestTop</CButton> -->
                                <CSelect hidden
                                  class="align-self-left"
                                  :value.sync="Prescript.drugsSearch"
                                  :options="drugs"
                                  @change="DisabledForPrescriptionOtherField"
                                  style="width: 100%"
                                >
                                </CSelect>
                                
                              
                              </CCol>

                              <CCol
                                col="5"
                                align="left"
                                style=""
                              >
                                <div>
                                  <CRow>
                                  <CLable hidden
                                    >Urgent &nbsp;</CLable>
                                  <CInputCheckbox hidden type="checkbox" id="Urgent"  :checked.sync="Prescript.urgentCheck"/>
                                  </CRow>
                                </div>
                              </CCol>

                              <CCol
                                col="2"
                                align="left"
                                style="margin-left: -180px"
                              >
                              
                              </CCol>

                              <CCol col="2">
                                <!-- <strong style="width: 12rem;">Gender : F</strong> -->
                              </CCol>
                              <CCol col="2">
                                <!-- <CButton color="info"  @input="InputSearch()"  style="width: 5rem; margin-left:4px;" class="BttPNs" @click="submit_of_health_checkup"> Save </CButton> -->
                              </CCol>
                            </CRow>

                            <CRow align="center"  style="margin-top: -15px">
                              <CCol col="5" style="margin-left: 0px">
                                
                                <CInput list="drugs" hidden
                                  placeholder="Input your search"
                                  v-model="Prescript.InputSearch"
                                  @input="onDrugSearch"
                                  @change="onDrugSelect"
                                />

                                <datalist id="drugs" loc>
                                  <option v-for="drug in InputSearchSelect" :key="drug.itemCode + ' ' + drug.value" :value="drug.value">{{drug.label}}</option>
                                </datalist>
                              </CCol>
                              <CCol col="2" style="margin-left: -25px">
                                <CInput hidden placeholder="Strenth" id="Strenth" v-model="Prescript.strenth"/>
                              </CCol>
                              <CCol
                                col="3"
                                align="left"
                                style="margin-left: -24px"
                              >
                                <!-- {{Prescript.selUnit}} -->
                                <CSelect hidden
                                  class="align-self-left"
                                  id="Unit"
                                  placeholder="Select Unit"
                                  :value.sync="Prescript.selUnit"
                                  :options="GetUnitList"
                                  style="width: 83%"
                                >
                                </CSelect>
                              </CCol>
                             &nbsp;&nbsp;&nbsp;
                              <CCol
                                col="3"
                                style="
                                  margin-left: -60px;
                                  padding: 0;
                                  max-width: 190px;
                                "
                              >
                                <!-- GetDosageList|oo| GetUnitList -->
                                <!-- {{Prescript.selFreque}} -->
                                <CSelect hidden
                                  class="align-self-right"
                                  id="Frequncy"
                                  placeholder="Select Frequency"
                                  :value.sync="Prescript.selFreque"
                                  :options="GetDosageList"
                                  style="width: 79%"
                                >
                                </CSelect>
                              </CCol>
                            </CRow>

                            <CRow align="center">
                              <CCol
                                col="5"
                                style="margin-left: 0px; margin-top: -13px"
                              >
                                <CInput hidden placeholder="Comment" v-model="Prescript.comment" />
                              </CCol>
                             
                              <CCol
                                col="4"
                                style="margin-left: -24px; margin-top: -13px"
                                align="left"
                              > 
                                <CSelect
                                  class="align-self-left"
                                  id="Duration" hidden
                                  placeholder="Duration"
                                  :value.sync="Prescript.week"
                                  :options="Dosageduration"
                                  style="width: 85%"
                                >
                                </CSelect>
                              </CCol>

                              <CCol
                                align="left"
                                col="2"
                                style="
                                  margin-left: -28px;
                                  margin-top: -13px;
                                  max-width: 190px;
                                "
                              >
                              <CButton color="info" hidden style="width: 8rem;" @click="AddForPrescription" disabled> {{ this.Prescript.index === null ? 'Add' : 'Update'}}</CButton>
                                
                              </CCol>
                      
                            </CRow>
                            <!-- <CCardHeader><strong> Urgent Prescriptions: </strong> </CCardHeader> -->
                            <!-- This Table for prescription  -->
                            <CRow style="margin-left: 0px; margin-top: -20px">
                              <CCol col="12">
                                <CLable><b>Prescriptions</b></CLable>
                                
                                <CDataTable
                                  header
                                  hover
                                  striped
                                  :items="DataforPrescriptonslist"
                                  :fields="DataforPrescriptonsField"
                                  :items-per-page="15"
                                  clickable-rows
                                  :active-page="activePage"
                                  @row-clicked="rowClicked"
                                  :pagination="{
                                    doubleArrows: false,
                                    align: 'center',
                                  }"
                                  @page-change="pageChange"
                                  scopedSlots
                                  sorter
                                  @loadstart="load"
                                >

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
                                    <i hidden style="cursor: no-drop;"
                                      ><CIcon
                                        name="cil-pencil"
                                        size="sm"
                                        class="ml-2 mb-2 text-primary"
                                    /></i>

                                    <i hidden style="cursor: no-drop;"
                                      ><CIcon
                                        name="cil-trash"
                                        size="sm"
                                        class="ml-2 mb-2 text-danger"
                                    /></i>
                                  </div>
                                </CButtonGroup>
                              </td>
                            </template>

                                </CDataTable>
                              </CCol>
                            </CRow>

                            <CRow style="margin-left: 0px; margin-top: -11px">
                              <CCol col="12">
                                <CLable><b>Urgent</b> Prescriptions</CLable>
                                <CDataTable
                                  header
                                  hover
                                  striped
                                  :items="UrgentPrescriptionDatalist"
                                  :fields="DataforPrescriptonsField"
                                  :items-per-page="5"
                                  clickable-rows
                                  :active-page="activePage"
                                  @row-clicked="rowClicked"
                                  :pagination="{
                                    doubleArrows: false,
                                    align: 'center',
                                  }"
                                  @page-change="pageChange"
                                  scopedSlots
                                  sorter
                                  @loadstart="load"
                                >
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
                                    <i hidden @click="editPrescription(item)"
                                      ><CIcon
                                        name="cil-pencil"
                                        size="sm"
                                        class="ml-2 mb-2 text-primary"
                                    /></i>

                                    <i hidden @click="deleteDiagnosisModal = true ; deletePrescriptionItem = item"
                                      ><CIcon
                                        name="cil-trash"
                                        size="sm"
                                        class="ml-2 mb-2 text-danger"
                                    /></i>
                                  </div>
                                </CButtonGroup>
                              </td>
                            </template>
                                </CDataTable>
                              </CCol>
                            </CRow>
                            <CRow>
                              <CCol col="12" align="center">
                                <CCard
                                  style="
                                    margin-left: 0px;
                                    margin-right: -15px;
                                  "
                                >
                                  <CRow>
                                   <CCol col="3" class="d-flex">
                                      <CLable
                                        style="
                                          margin-top: 5px;
                                          margin-left: 20px;
                                        "
                                        >Clinic</CLable
                                      >
                                      <!-- localStorage.getItem("Location") -->
                                      <CSelect
                                        class="col align-self-center"
                                        :value.sync="Prescript.clinic"
                                        :options="ClinicType"
                                        placeholder="Clinic"
                                        style="margin-left: -10px"
                                      >
                                      </CSelect>
                                    </CCol>
                                    <CCol
                                      col="5"
                                      class="d-flex"
                                      style="margin-left: -35px"
                                    >
                                      <CLable
                                        style="
                                          margin-top: 5px;
                                          margin-left: 10px;
                                        "
                                        >Next Clinic in</CLable
                                      >
                                      <CSelect
                                        class="col align-self-center"
                                        style="
                                          margin-left: -10px;
                                          max-width: 90px;
                                        "
                                        :value.sync="header.nextClinicInValue"
                                        :options="Days"
                                        placeholder="30"
                                      ></CSelect>
                                      <CSelect
                                        class="col align-self-center"
                                        style="margin-left: -25px; width: 20%"
                                        :value.sync="header.nextClinicInUnit"
                                        :options="NextClinicIn"
                                        placeholder="Next Clinic In"
                                      >
                                      </CSelect>
                                    </CCol>

                                    <CCol
                                      col="5"
                                      class="d-flex"
                                      style="margin-left: -50px"
                                    >
                                      <CLable
                                        style="
                                          margin-top: 5px;
                                          margin-left: 20px;
                                          margin-right: 6px;
                                          padding-left: 20px;
                                        "
                                        >Annual clinic visit</CLable
                                      >
                                      <!-- <CInput type="date" /> -->
                                      <CSelect
                                        class="col align-self-center"
                                        style="
                                          margin-left: -25px;
                                          padding-left: 20px;
                                        "
                                        :value.sync="header.annualVisit"
                                        :options="AnnualClinicvisit"
                                        placeholder="---"
                                      >
                                      </CSelect>
                                    </CCol>
                                  </CRow>

                                  <CRow>
                                    <CCol style="margin-top: -14px">
                                      <CTextarea
                                        placeholder="Clinic service"
                                        v-model="Prescript.presentcomplaint"
                                        :is-valid="presentcomplaintIsValid"
                                      />
                                    </CCol>
                                  </CRow>
                                  <CRow>
                                    <CCol
                                      style="
                                        margin-top: -13px;
                                        margin-bottom: 2px;
                                      "
                                    >
                                      <!-- <CButton
                                        color="behance"
                                        style="width: 5rem; margin-left: 4px"
                                      >
                                        Save
                                      </CButton>
                                      <CButton
                                        color="reddit"
                                        style="width: 5rem; margin-left: 4px"
                                      >
                                        Clear
                                      </CButton> -->
                                    </CCol>
                                  </CRow>
                                </CCard>
                              </CCol>
                            </CRow>
                          </CCardBody>
                          <!-- </CCard> -->
                    </CTabPane>
                    </CTabContent>
                  </CTab>
                  <CTab title="Comments">
                    <CTabContent style="margin-top: 50px; padding: 0px">
                      <CTabPane
                          class=""
                          style="padding-left: 26px; margin-bottom: 30px"
                        >
                          <CCardBody
                            style="margin-left: -9px; margin-top: -20px"
                          >
                            <!-- </CCol> -->
                          </CCardBody>
                          <CRow style="margin-top: -20px;padding:5px">
                            <CCol col="9">
                              <form>
                                <!-- * Here the firaz said to unable to validation  tue-08-02-2022-->
                                <CTextarea
                                  label="Comment(Print)"
                                  valid-feedback="Input is valid."
                                  invalid-feedback="Please provide comment(Print)."
                                  value="Valid value"
                                  v-model="header.comment_print"
                                  rows="8"
                                  cols="50"
                                />
                                <CTextarea
                                  label="Comment(Internal)"
                                  v-model="header.comment_internal"
                                  rows="8"
                                  cols="50"
                                />
                              </form>
                            </CCol> 
                            <!-- &nbsp;&nbsp; -->
                            <CCol col="3" style="margin-left: -25px">
                              <!-- <h6 @click="modaltest()">Click My Tesk</h6> -->
                              <CCard
                                style="margin-right: -30px; margin-top: 5px"
                              >
                                <div class="custom-file">
                                  <div @click="uploadImage" hidden>
                                    <input hidden
                                      type=""
                                      class="custom-file-input"
                                      style="cursor: pointer"
                                      id="customFile"
                                      @click="uploadImage"
                                    />
                                    <label hidden
                                      class="custom-file-label"
                                      for="customFile"
                                      >Choose file</label
                                    >
                                  </div>
                                </div>
                                <div v-for="data in ImageOfDisease" :key="data">
                                  <CRow
                                    class="rounded shadow bg-body border-info"
                                    style="
                                      margin-left: -1px;
                                      margin-right: -1px;
                                      margin-top: 2.5px;
                                      margin-bottom: 2px;
                                      cursor: pointer;
                                    "
                                    @click="imageload()"
                                  >
                                    <!-- working fine temp hold -->

                                    <CCol style="margin-left: -14px">
                                      <img
                                        v-bind:src="data.cNameofImage"
                                        class="shadow rounded"
                                        style="
                                          max-width: 36px;
                                          max-height: 36px;
                                          margin: 0.5px;
                                        "
                                      />
                                    </CCol>
                                    <CCol style="margin-left: -125px">
                                      <CRow class="d-flex">
                                        <CCol>
                                          <p
                                            class="text-primary"
                                            style="
                                              font-size: small;
                                              word-spacing: -1px;
                                              margin: 0;
                                              padding: 0;
                                              line-height: 0.9;
                                            "
                                          >
                                            {{ data.cNameofdisease }}
                                          </p>
                                        </CCol>
                                      </CRow>
                                      <CRow
                                        class="d-flex"
                                        style="
                                          margin-left: 1px;
                                          margin-right: -18px;
                                          margin-bottom: -22px;
                                        "
                                      >
                                        <CCol
                                          size="sm"
                                          align="left"
                                          class="d-flex"
                                          style="margin-left: -14px"
                                        >
                                          <p style="font-size: x-small">
                                            {{ data.UserId }}
                                          </p></CCol
                                        >
                                        <CCol
                                          size="sm"
                                          align="right"
                                          style="max-width: 50%"
                                          class="d-flex text-primary"
                                        >
                                          <p
                                            style="font-size: xx-small"
                                            class="text-success"
                                          >
                                            {{
                                              new Date(
                                                data.CreateDate
                                              ).toLocaleDateString()
                                            }}
                                          </p></CCol
                                        >
                                      </CRow>
                                    </CCol>
                                  </CRow>
                                </div>
                              </CCard>
                            </CCol>
                          </CRow>
                        </CTabPane>
                    </CTabContent>
                  </CTab>

   <!-- ============================================================================================================================ -->
   <!--  REFERRAL  TAP -->
   <!--  Author: Ravindra Kumara -->
   <!--  Create date: 2022 March -->
   <!--  DISCRIPTION: "" -->
   <!--  Modify date: 2022 March 23-->
   <!-- ============================================================================================================================ -->

                  <CTab title="Referral Form">
                    <CTabContent style="margin-top: 50px; padding: 0px">
                     <CTabPane v-if="tabIndex === 4">
                        <CRow>
                            <CCol class="" style="margin-left:2px;">
                                <CInput placeholder="DOCTOR NAME" v-model="ReferralForm.Doctor" disabled/>
                            </CCol>
                        </CRow>
                          <CForm style="height: 400px">
                            <quill-editor
                              v-model="content"
                              :content="ReferralForm.Letter"
                              style="height: 350px"
                            />
                          </CForm>
                          <CRow class="" style="margin-top: 20px">
                            <CCol class="d-flex">
                              <CButton hidden
                                color="facebook"
                                style="
                                  width: 10rem;
                                  height: 2.3rem;
                                  margin-right: 30px;
                                "
                                @click="CallRefNameTempModel"
                              >
                                Save as Template
                              </CButton>

                              <CButton hidden
                                color="tumblr"
                                style="
                                  width: 8rem;
                                  margin-right: 30px;
                                  height: 2.3rem;
                                "
                                @click="SaveasReferralform"
                              >
                                Save as Referral
                              </CButton>

                              <CButton
                                color="yahoo"
                                style="width: 8rem; height: 2.3rem"
                                align="right" hidden
                                @click="loadTemplate"
                              >
                                Load Template
                              </CButton>

                             <CSelect
                                class="col align-self-right"
                                :value.sync="LetterID" disabled
                                :options="ReferralLetter"
                                placeholder="ReferralID"
                                @click="GetPatientReferral(item)"
                                style="
                                  margin-left: -10px;
                                  margin-right: -15px;
                                  width: 20%;
                                ">
                              </CSelect>
                            </CCol>
                          </CRow>
                               
                         

                          <div class="align_set">
                            &nbsp; &nbsp; &nbsp;
                            <!-- <CButton color="info" style="width: 8rem;" @click="reset_referral" disabled> Clear </CButton> -->
                          </div>
                        </CTabPane>
                    </CTabContent>
                  </CTab>


                  <CTab title="Initial screening">
                    <CTabContent style="margin-top: 50px; padding: 0px">
                      <CTabPane>
                          <form style="margin-left:10px">
                            <CTextarea
                              label="Presenting Complaint"
                              v-model="initScreen.presentcomplaint"
                              :is-valid="presentcomplaintIsValid"
                            />
                            <!-- <p v-if="!" class="error-showIS1"></p> -->
                            <CRow md="12">
                              <CCol md="6">
                                <CTextarea
                                  label="Past medical / Surgical History"
                                  v-model="initScreen.past_medi_or_surgicalHx"
                                  :is-valid="past_medi_or_surgicalHxIsValid"
                                />
                                <!-- <p v-if="!" class="error-showIS2">! Past medical /surgicalHistory</p> -->
                              </CCol>
                              <CCol md="6">
                                <CTextarea
                                  label="Obstretic / Menstrual History"
                                  v-model="initScreen.obstretic_or_menstrualHx"
                                  :is-valid="obstretic_or_menstrualHxIsValid"
                                />
                                <!-- <p v-if="!" class="error-showIS2">! Obstretic/menstrual History</p> -->
                              </CCol>
                            </CRow>
                            <CRow md="12">
                              <CCol md="6">
                                <CTextarea
                                  label="Medication summary"
                                  v-model="initScreen.medication_summary"
                                  :is-valid="medication_summaryIsValid"
                                />
                                <!-- <p v-if="!" class="error-showIS2">! Past medical /surgical History</p> -->
                              </CCol>
                              <CCol md="6">
                                <CTextarea
                                  label="Family History"
                                  v-model="initScreen.familtyHistory"
                                  :is-valid="familtyHistoryIsValid"
                                />
                                <!-- <p v-if="!" class="error-showIS2">! Obstretic/menstrual History</p> -->
                              </CCol>
                            </CRow>

                            <!-- <p v-if="!" class="error-showIS1">! Medication summary</p> -->

                            <!-- <p v-if="!" class="error-showIS1">! Familty History</p> -->
                            <CRow md="12">
                              <CCol md="6">
                                <CTextarea
                                  label="Allergic History"
                                  v-model="initScreen.allergicHx"
                                  :is-valid="allergicHxIsValid"
                                />
                                <!-- <p v-if="!" class="error-showIS2">! Allergic History</p> -->
                              </CCol>
                              <CCol md="6">
                                <CTextarea
                                  label="Social History"
                                  v-model="initScreen.socialHx"
                                  :is-valid="socialHxIsValid"
                                />
                                <!-- <p v-if="!" class="error-showIS2">! Social History</p> -->
                              </CCol>
                            </CRow>

                            <CRow md="12">
                              <CCol md="6">
                                <CTextarea
                                  label="General Examination"
                                  v-model="initScreen.generalExamination"
                                  :is-valid="generalExaminationIsValid"
                                />
                                <!-- <p v-if="!" class="error-showIS2">! General Examination</p> -->
                              </CCol>

                              <CCol md="3">
                                <CFormLabel
                                  for="inputPassword3"
                                  class="col-sm-2 col-form-label"
                                  >Smoking</CFormLabel
                                >
                                <div style="display: flex" align="center">
                                  <!-- <div class="margin_set">

                      </div> -->

                                  <CInput
                                    type="radio"
                                    style="margin-left: 5px"
                                    name="fav_language"
                                    value="HTML"
                                    label="Precent"
                                  />
                                  <CInput
                                    type="radio"
                                    style="margin-left: 5px"
                                    name="fav_language"
                                    value="HTML"
                                    label="Past"
                                  />
                                  <CInput
                                    type="radio"
                                    style="margin-left: 5px"
                                    name="fav_language"
                                    value="HTML"
                                    label="None"
                                  />
                                </div>
                                <!-- <p v-if="!socialHxIsValid" class="error-showIS2">! Smorking</p> -->
                              </CCol>
                              <CCol md="3">
                                <CFormLabel
                                  for="inputPassword3"
                                  class="col-sm-2 col-form-label"
                                  >Alcohol
                                  <div style="display: flex" align="center">
                                    <CInput
                                      type="radio"
                                      style="margin-left: 5px"
                                      name="fav_language"
                                      value="HTML"
                                      label="Precent"
                                    />
                                    <CInput
                                      type="radio"
                                      style="margin-left: 5px"
                                      name="fav_language"
                                      value="HTML"
                                      label="Past"
                                    />
                                    <CInput
                                      type="radio"
                                      style="margin-left: 5px"
                                      name="fav_language"
                                      value="HTML"
                                      label="None"
                                    />
                                  </div>
                                </CFormLabel>
                                <!-- <p v-if="!socialHxIsValid" class="error-showIS2">! Alcohol</p> -->
                              </CCol>
                            </CRow>

                            <!-- <CTextarea label="General Examination" /> -->

                            <CRow md="12">
                              <CCol md="6">
                                <CTextarea
                                  label="FGS"
                                  v-model="initScreen.fGS"
                                  :is-valid="fGSIsValid"
                                />
                                <!-- <p v-if="!fGSIsValid" class="error-showIS1" >! FGS</p> -->
                              </CCol>
                              <CCol md="6">
                                <CFormLabel for="formFileSm"
                                  >FGS Image</CFormLabel
                                >
                                <CInput
                                  type="file"
                                  size="sm"
                                  id="formFileSm"
                                  v-model="initScreen.fGS_Image"
                                  :is-valid="fGS_ImageIsValid"
                                />
                                <!-- <p v-if="!fGS_ImageIsValid" class="error-showIS1">! FGS Image</p> -->
                              </CCol>
                            </CRow>
                            <CRow md="12">
                              <CCol md="6">
                                <CTextarea
                                  label="Thyroid Examination"
                                  v-model="initScreen.Thyroid_Examination"
                                  :is-valid="Thyroid_ExaminationIsValid"
                                />
                                <!-- <p v-if="!Thyroid_ExaminationIsValid" class="error-showIS1">! Thyroid Examination</p> -->
                              </CCol>
                              <CCol md="6">
                                <CTextarea
                                  label="Eyes"
                                  v-model="initScreen.eyes"
                                  :is-valid="eyesIsValid"
                                />
                                <!-- <p v-if="!eyesIsValid" class="error-showIS1">! Eyes</p> -->
                              </CCol>
                            </CRow>

                            <CRow md="12">
                              <CCol md="6">
                                <CTextarea
                                  label="RS"
                                  v-model="initScreen.rS"
                                  :is-valid="rSIsValid"
                                />
                                <!-- <p v-if="!rSIsValid" class="error-showIS1">! RS</p> -->
                              </CCol>

                              <CCol md="6">
                                <CTextarea
                                  label="ABD"
                                  v-model="initScreen.aBD"
                                  :is-valid="aBDIsValid"
                                />
                                <!-- <p v-if="!aBDIsValid" class="error-showIS1">! ABD</p> -->
                              </CCol>
                            </CRow>

                            <CRow md="12">
                              <CCol md="6">
                                <CTextarea
                                  label="CNS"
                                  v-model="initScreen.cNS"
                                  :is-valid="cNSIsValid"
                                />
                                <!-- <p v-if="!cNSIsValid" class="error-showIS1">! CNS</p> -->
                              </CCol>
                              <CCol md="6">
                                <CTextarea
                                  label="CVS"
                                  v-model="initScreen.cVS"
                                  :is-valid="cVSIsValid"
                                />
                                <!-- <p v-if="!cVSIsValid" class="error-showIS1">! CVS</p> -->
                              </CCol>
                            </CRow>

                            <!-- </CForm>  -->
                          </form>
                          <!-- <div class="align_set">
                            <CButton
                              color="info"
                              style="width: 8rem"
                              @click="Submit_of_Initial_Screen"
                            >
                              Save
                            </CButton>
                            &nbsp;
                            <CButton color="info" style="width: 8rem">
                              Clear
                            </CButton>
                          </div> -->

                          <!-- </CCard> -->

                          <!-- </CCol> -->
                          <!-- </CRow> -->
                        </CTabPane>
                    </CTabContent>
                  </CTab>
                  <CTab title="Checked list">
                    <CTabContent style="margin-top: 50px; padding: 0px">
                    <CTabPane >

                          <CRow md="9" v-for="item in checkedList" style="margin-left:5px" :key="item.ItemCode" align="left">
                            <CCol col="4">
                              <CInputCheckbox
                                type="checkbox"
                                :label="item.Description"
                                class="input_checkCL"
                                valid-feedback="Thank you :)"
                                invalid-feedback="Please provide at least 4 characters."
                                :is-valid="validator"
                                :checked.sync="item.value"
                              />
                            </CCol>
                          </CRow>


                          <CForm> </CForm>
                          <div class="align_set" style="margin-left:10px">
                            <CButton hidden
                              color="behance"
                              style="width: 8rem"
                              @click="onCheckListSave"
                            >
                              Save
                            </CButton>
                            &nbsp;
                            <CButton hidden color="twitter" style="width: 8rem">
                              Clear
                            </CButton>
                          </div>
                          <CDataTable
                            class=""
                            style=" margin-left:10px;
                              margine-bottom: 10px;
                              margine-top: 60px;
                              padding-top: 4px;
                            "
                            header
                            hover
                            striped
                            :items="checkedGrid"
                            :fields="checkedGridFields"
                            :items-per-page="10"
                            clickable-rows
                            :active-page="activePage"
                            @row-clicked="rowClicked"
                            :pagination="{
                              doubleArrows: false,
                              align: 'center',
                            }"
                            @page-change="pageChange"
                            scopedSlots
                            sorter
                            @loadstart="load"
                          >
                          </CDataTable>
                        </CTabPane>
                        </CTabContent>
                  </CTab>
                  <CTab title="Pregnany Diabetes" hidden>
                    <CTabContent style="margin-top: 50px; padding: 0px">
                     <CTabPane >
                          <CForm>
                            <CInput
                              label="ID:"
                              valid-feedback="Input is valid."
                              invalid-feedback="Please Enter The ID."
                              v-model="PregnancyDiabetes.id"
                              :is-valid="Idvalidator"
                            />

                            <CRow md="12">
                              <CCol md="4">
                                <CInput
                                  label="Date:"
                                  valid-feedback="Input is valid."
                                  invalid-feedback="Please Enter The DATE."
                                  v-model="PregnancyDiabetes.date"
                                  :is-valid="Datevalidator"
                                />
                              </CCol>
                              <CCol md="4">
                                <CInput
                                  label="POA"
                                  valid-feedback="Input is valid."
                                  invalid-feedback="Please Enter The POA."
                                  v-model="PregnancyDiabetes.pOA"
                                  :is-valid="POAvalidator"
                                />
                              </CCol>
                              <CCol md="4">
                                <CInput
                                  label="Obstretic"
                                  valid-feedback="Input is valid."
                                  invalid-feedback="Please Enter The Obstretic."
                                  v-model="PregnancyDiabetes.obstretic"
                                  :is-valid="Obstreticvalidator"
                                />
                              </CCol>
                            </CRow>
                            <table style="width: 100%">
                              <tr>
                                <th style="width: 20%">R% regime</th>
                                <th>Meals</th>
                                <th>Before</th>
                                <th>After</th>
                                <th style="width: 20%">Adjusted R% regime</th>
                              </tr>
                              <tr>
                                <td>
                                  <CTextarea
                                    valid-feedback="Input is valid."
                                    invalid-feedback="Please Enter The R % regime."
                                    v-model="PregnancyDiabetes.rregime"
                                    :is-valid="Rregimevalidator"
                                  />
                                </td>
                                <td>
                                  <label>Breakfast</label>
                                  <br />
                                  <br />
                                  <br />

                                  <label>Lunch</label>
                                  <br />
                                  <br />
                                  <br />
                                  <br />
                                  <label>Dinner</label>
                                </td>
                                <td>
                                  <CInput
                                    valid-feedback="Input is valid."
                                    invalid-feedback="Please Enter The BreakFast Before."
                                    v-model="PregnancyDiabetes.bfbvalue"
                                    :is-valid="Bfbvaluevalidator"
                                  />
                                  <CInput
                                    valid-feedback="Input is valid."
                                    invalid-feedback="Please Enter The Lunch Before."
                                    v-model="PregnancyDiabetes.lbvalue"
                                    :is-valid="LBvaluevalidator"
                                  />
                                  <CInput
                                    valid-feedback="Input is valid."
                                    invalid-feedback="Please Enter The Dinner Before."
                                    v-model="PregnancyDiabetes.dbvalue"
                                    :is-valid="Dbvaluevalidator"
                                  />
                                </td>
                                <td>
                                  <CInput
                                    valid-feedback="Input is valid."
                                    invalid-feedback="Please Enter The BreakFast After."
                                    v-model="PregnancyDiabetes.bfavalue"
                                    :is-valid="BFAvaluevalidator"
                                  />
                                  <CInput
                                    valid-feedback="Input is valid."
                                    invalid-feedback="Please Enter The Lunch After."
                                    v-model="PregnancyDiabetes.lavalue"
                                    :is-valid="LAvaluevalidator"
                                  />
                                  <CInput
                                    valid-feedback="Input is valid."
                                    invalid-feedback="Please Enter The Dinner After."
                                    v-model="PregnancyDiabetes.davalue"
                                    :is-valid="DAvalidator"
                                  />
                                </td>
                                <td>
                                  <CTextarea
                                    label="Comment (Print)"
                                    valid-feedback="Input is valid."
                                    invalid-feedback="Please Enter The Comment Print."
                                    v-model="PregnancyDiabetes.aregime"
                                    :is-valid="Aregimevalidator"
                                  />
                                </td>
                              </tr>
                            </table>
                            <CTextarea
                              label="Notes"
                              valid-feedback="Input is valid."
                              invalid-feedback="Please Enter The Notes"
                              v-model="PregnancyDiabetes.rnotes"
                              :is-valid="Rnotevalidator"
                            />
                          </CForm>
                          <div class="align_set">
                            <CButton
                              color="info"
                              style="width: 8rem"
                              @click="Submit_of_Pregnancy_diabetes"
                            >
                              Save
                            </CButton>
                            &nbsp;
                            <CButton color="info" style="width: 8rem">
                              Clear
                            </CButton>
                          </div>
                        </CTabPane>
                        </CTabContent>
                  </CTab>
                  <CTab title="Footcare" hidden>
                   <CTabContent style="margin-top: 50px; padding: 0px">
                   <CTabPane> Footcare </CTabPane>
                    </CTabContent>
                  </CTab>
                </CTabs>
              </CCol>

              <CCol col="2">
                <CCardBody>
                  <CCol
                    class="d-flex"
                    style="
                      margin-left: -38px;
                      margin-right: -5px;
                      padding: 0px;
                      margin-top: -10px;
                    "
                  >
                    <CButton hidden color="success" class="m-1" disabled
                      >Save
                    </CButton>
                    <CDropdown
                      toggler-text="Print"
                      style="margin-top: 4px"
                      color="yahoo"
                    >
                      <CDropdownItem disabled>Footcare</CDropdownItem>
                      <CDropdownItem>Initial visit</CDropdownItem>
                      <CDropdownItem>Referral form</CDropdownItem>
                      <CDropdownItem @click="onShowSummaryPDF" >Summary</CDropdownItem>
                      <CDropdownItem>Urgent prescriptions</CDropdownItem>
                    </CDropdown>
                    <!-- <CButton color="success"  class="m-1" @click="Prescription()">Print</CButton> -->
                  </CCol>
                </CCardBody>
              </CCol>
            </CRow>

            

          </CCard>
        </CCol>
      </CRow>
      <!-- All the Content are ENDED -->
    </div>
  </div>
</template>

<script>
import Initial_screening from "../Initial_Screening/Initial_screening.vue";
import { DateTime } from 'luxon'
import Vue from "vue";
import VueQuillEditor from "vue-quill-editor";
import "quill/dist/quill.core.css";
import "quill/dist/quill.snow.css";
import "quill/dist/quill.bubble.css";
// import Cropper from "cropper.js"
Vue.use(VueQuillEditor);

export default {
  inject: ["api"],
  components: { Initial_screening },

  data() {
    return {
      id: null,
      err: "",
      // here for Referral 
      ReferraltemplateName:false,
      ReferralTemp:null,
      ReferralLetterTemplateArray:[],
      //
      alertboxSuccess: false,
      alertboxFailed: false,
      deletePrescriptionModal : false,
      deletePrescriptionItem: null,
      deleteDiagnosisModal : false,
      deleteDiagnosisItem: null,
      ClinicSession: null,
      NewSession:[{"label": "New","value":"New"}],
      ClinicWiseSession: null,
      DocumentHeaderNotes: null,
      problemList: null,
      GetDocumentHeaderImagelist: null,
      GetDosageList: null,
      GetUnitList: null,
      Profile: null,
      BodyMI: [],
      labAndDrug: null,
      DP: localStorage.getItem("DP"),
      ImageOfDisease: null,
      ProfileImage: null,
      imagePreviewFordisease:null,
      Photo: null,
       // PRESCRIPTION MAIN ARRAY
      DataforPrescriptons: [],

      // FOR PRESCRIPTION TABLE
      DataforPrescriptonslist:[],

      // URGENT PRESCRIPTION TABLE  
      UrgentPrescriptionDatalist:[],
      // very Importan .PulseRate
      //
      diagnosticList: [],
      //Problem Input

      ProblemDescriptionValue: null,
      ProblemNote: null,
      ProblemPriority: null,
      ProblemStatus: null,

      ReferralLetter: null,
      ReferralLetterTemplate: null,
      ReferralLetterTemplateField: [
        {
          filter: false,
          sorter: false,
          key: "TemplateName",
          label: "Template Name",
        },
        {
          filter: false,
          sorter: false,
          key: "vcConsultantName",
          label: "Doctor",
        },

        {
          filter: false,
          sorter: false,
          key: "call",
          label: "",
          _props: { color: "primary", className: "fw-semibold" },
          _style: { width: "1%" },
        },
      ],
      //
      drugs: [
        { lable: "Drugs", value: "Drugs" },
        { lable: "Lab Investication", value: "Lab" },
        { lable: "Other Investication", value: "Other" },
      ],
      drugsSearch: null,
      AnnualClinicvisit: [
        { lable: "January", value: "January" },
        { lable: "February", value: "February" },
        { lable: "March", value: "March" },
        { lable: "April", value: "April" },
        { lable: "May", value: "May" },
        { lable: "June", value: "June" },
        { lable: "July", value: "July" },
        { lable: "August", value: "August" },
        { lable: "September", value: "September" },
        { lable: "October", value: "October" },
        { lable: "November", value: "November" },
        { lable: "December", value: "December" },
      ],
      Days: [
        { label: "1", value: "1" },
        { label: "2", value: "2" },
        { label: "3", value: "3" },
        { label: "4", value: "4" },
        { label: "5", value: "5" },
        { label: "6", value: "6" },
        { label: "7", value: "7" },
        { label: "8", value: "8" },
        { label: "9", value: "9" },
        { label: "10", value: "10" },
        { label: "11", value: "11" },
        { label: "12", value: "12" },
        { label: "13", value: "13" },
        { label: "14", value: "14" },
        { label: "15", value: "15" },
        { label: "16", value: "16" },
        { label: "17", value: "17" },
        { label: "18", value: "18" },
        { label: "19", value: "19" },
        { label: "20", value: "20" },
        { label: "21", value: "21" },
        { label: "22", value: "22" },
        { label: "23", value: "23" },
        { label: "24", value: "24" },
        { label: "25", value: "25" },
        { label: "26", value: "26" },
        { label: "27", value: "27" },
        { label: "28", value: "28" },
        { label: "29", value: "29" },
        { label: "30", value: "30" },
        { label: "31", value: "31" },
      ],

      NextClinicIn: [
        { label: "Years", value: "Years" },
        { label: "Days", value: "Days" },
        { label: "Weeks", value: "Weeks" },
        { label: "Months", value: "Months" },
      ],
      Usertype: "endocrine",
      InputSearchSelect:null,
      tabIndex: 0,
      loadtemplate: false,
      mytext: false,
      warningModal: false,
      imagePreview: "/img/avatars/blank.png",
      AltImage: "Image.jpg",
      imageModal: false,
      //
      patient: null,
      InputSearch: null,
      //
      content: null,
      // for health check
      header: {
        bp: null,
        pulse_rate: null,
        weight: 0,
        height: 0,
        bmi: 0,
        classificationValue: null,
        status: null,
        status_com: null,
        comment_print: null,
        comment_internal: null,
        nextClinicIn: null,
        nextClinicInValue: null,
        nextClinicInUnit: null,
        annualVisit: null,
      },
      ClinicType:[{label:"Medical Clinic",value:"MC"},{label:"Endocrine",value:'D&E'}],
      //diagnosis
      diagnosis: {
        index: null,
        ProblemDescriptionValue: null,
        ProblemNote: null,
        ProblemPriority: 1,
        ProblemStatus: "Active",
      },
      //    Prescription
      Prescript: {
        index: null,
        drugsSearch: "Drugs",
        urgentCheck: null,
        selectedDrug: null,
        InputSearch: null,
        strenth: null,
        selUnit: null,
        selFreque: null,
        comment: null,
        week: null,
        clinic: null,
        nexClinic: null,
        nexClinicIn: null,
        annualVisit: null,
        presentcomplaint: null,
        //  getunit:null,
      },
      SearchValue: null,
      // Initial Screen
      initScreen: {
        id: 0,
        presentcomplaint: null,
        past_medi_or_surgicalHx: null,
        obstretic_or_menstrualHx: null,
        medication_summary: null,
        familtyHistory: null,
        allergicHx: null,
        socialHx: null,
        generalExamination: null,
        smorking: null,
        alcohol: null,
        fGS: null,
        fGS_Image: null,
        Thyroid_Examination: null,
        eyes: null,
        rS: null,
        aBD: null,
        cNS: null,
        cVS: null,
      },

      classificationDropdown: [
        { label: "Diabetes", value: "Diabetes" },
        { label: "Thyroid", value: "Thyroid" },
        { label: "Pregnancy Diabetes", value: "Pregnancy_Diabetes" },
        { label: "Endocrine", value: "Endocrine" },
      ],

      statusDropdown: [
        { label: "Regular", value: "Regular" },
        { label: "Yearly Only", value: "Yearly_Only" },
        { label: "Transfer", value: "Transfer" },
        { label: "Pregnancy Diabetes", value: "Pregnancy_diabetes" },
        { label: "Discharge", value: "discharge" },
      ],

      labvalue: null,
      // Check list
      CheckList: {
        annualReview: null,
        footExamination: null,
        eyesReview: null,
        dietatyAdvice: null,
        insulinInjection: null,
        sMBG: null,
        sickDayRules: null,
        aTDsideeffects: null,
        bisphoshonateAdvice: null,
      },

      //Referrals
      ReferralForm: {
        referral_id: null,
        rdoctor: null,
        rnotes: null,
      },
      //PregnancyDiabetes

      // PrescriptionField:[{
      //     filter: false,
      //     sorter: false,
      //     key: "",
      //     label: "Category",
      //   },
      //   {
      //     filter: false,
      //     sorter: false,
      //     key: "",
      //     label: "Description",
      //   },
      //   {
      //     filter: false,
      //     sorter: false,
      //     key: "",
      //     label: "Dosage",
      //   }],
      // Prescription:null,
      PregnancyDiabetes: {
        id: null,
        date: null,
        pOA: null,
        obstretic: null,
        bfbvalue: null,
        lbvalue: null,
        dbvalue: null,
        bfavalue: null,
        lavalue: null,
        davalue: null,
        rregime: null,
        aregime: null,
        rnotes: null,
      },
      errors: [],
      Alertbox_success: false,
      Dosageduration: [
         {value: "1 Day"  , label: "1 Day" }, 
        {value: "2 Days" , abel: "2 Days" },
        {value: "3 Days" , label: "3 Days"},
        {value: "4 Days" , label: "4 Days"},
        {value: "5 Days" , label: "5 Days"},
        {value: "6 Days" , label: "6 Days"}, 
        {value: "7 Days" , label: "7 Days"}, 
        {value: "8 Days" , label: "8 Days"}, 
        {value: "9 Days" , label: "9 Days" },
        {value: "10 Days", label: "10 Days"},
        {value: "Daily" , label: "Daily" },
        {value: "1 Week" , label: "1 Week"},
        {value: "2 Weeks" ,label: "2 Weeks" },
        {value: "3 Weeks" ,label: "3 Weeks" },
        {value: "6 Weeks" , label: "6 Weeks"},
        {value: "1 Month" ,label: "1 Month" },
        {value: "2 Months" ,label: "2 Months" },
        {value: "3 Months" , label: "3 Months"},
        {value: "4 Months", label: "4 Months"},
        {value: "5 Months" ,label: "5 Months"},
        {value: "6 Months" , label: "6 Months"},
      ],
      annualVisit: null,
      BMMI: null,
      TemplateFields: [
        {
          filter: false,
          sorter: false,
          key: "TempID",
          label: "Template ID",
        },
        {
          filter: false,
          sorter: false,
          key: "TempName",
          label: "Template Name",
        },
        {
          filter: false,
          sorter: false,
          key: "DocID",
          label: "Doctor ID",
        },
        {
          filter: false,
          sorter: false,
          key: "Letter",
          label: "Letter",
        },
        {
          filter: false,
          sorter: false,
          key: "call",
          label: "",
          _props: { color: "primary", className: "fw-semibold" },
          _style: { width: "1%" },
        },
      ],
      Template: [
        {
          id: 1,
          TempName: "eye Ref",
          TempID: "00123",
          DocID: "Dr.Dhayanandana",
          Letter:
            "Dear Sir,Please be kind enough to see this 67 year old patient who is followed up in our clinic for type 2 DM with moderate control and C/O Left side lower groin pain mostly when lifting a heavy weight. ?? Ingunal Hernia. There are no other sinister signs or symptoms",
        },
        {
          id: 2,
          TempName: "hair Ref",
          TempID: "00124",
          DocID: "Dr.Dhayanandana",
          Letter:
            "Dear Sir,Please be kind enough to see this 67 year old patient who is followed up in our clinic for type 2 DM with moderate control and C/O Left side lower groin pain mostly when lifting a heavy weight. ?? Ingunal Hernia. There are no other sinister signs or symptoms",
        },
        {
          id: 3,
          TempName: "leg Ref",
          TempID: "00125",
          DocID: "Dr.Sunil",
          Letter:
            "Dear Sir,Please be kind enough to see this 67 year old patient who is followed up in our clinic for type 2 DM with moderate control and C/O Left side lower groin pain mostly when lifting a heavy weight. ?? Ingunal Hernia. There are no other sinister signs or symptoms",
        },
      ],
      checkedGridFields: [
        { filter: false, sorter: false, key: "DocumentID", label: "DocumentID" },
      ],
      DataforPrescriptonsField: [
        {
          filter: false,
          sorter: false,
          key: "ItemType",
          label: "Category",
          _style: { width: "14%" },
        },
        {
          filter: false,
          sorter: false,
          key: "Description",
          label: "Descriptions",
          _style: { width: "14%" },
        },
        {
          filter: false,
          sorter: false,
          key: "Dosage",
          label: "Dosage",
          _style: { width: "14%" },
        },
        {
          filter: false,
          sorter: false,
          key: "Comment",
          label: "Comment",
          _style: { width: "14%" },
        },
        {
          key: "show_details",
          label: "Action",
          _style: { width: "1%" },
          filter: false,
          sorter: false,
          _props: { color: "primary", className: "fw-semibold mr-20" },
        },
      ],
      problemListFields: [
        {
          filter: false,
          sorter: false,
          // key: "ProblemDescription"
          // key: "Clinic", ,"Clinic"
          key: "ProblemDescription",
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
        {
          filter: false,
          sorter: false,
          key: "Priority",
          label: "Priority",
          _style: { width: "10%" },
        },
        {
          key: "show_details",
          label: "Action",
          _style: { width: "1%" },
          filter: false,
          sorter: false,
          _props: { color: "primary", className: "fw-semibold mr-20" },
        },
      ],
      checkedGrid: null,
      checkedList: [],
      LabInvestfields: [
        {
          filter: false,
          sorter: false,
          // key: "ProblemDescription"
          // key: "Clinic", ,"Clinic"
          key: "parameters",
          label: "Test parameters",
          _style: { width: "14%" },
        },
        {
          filter: false,
          sorter: false,
          key: "inputfield",
          label: "Today",
          _style: { width: "5%" },
        },
        {
          filter: false,
          sorter: false,
          key: "date",
          label: "03/02/21",
          _style: { width: "10%" },
        },
        {
          key: "date",
          label: "03/05/21",
          _style: { width: "1%" },
          filter: false,
          sorter: false,
          _props: { color: "primary", className: "fw-semibold mr-20" },
        },

        {
          key: "date",
          label: "03/06/21",
          _style: { width: "1%" },
          filter: false,
          sorter: false,
          _props: { color: "primary", className: "fw-semibold mr-20" },
        },
      ],

      LabInvestData: [
        { id: 1, parameters: "WPC", date: "640" },
        { id: 2, parameters: "RPC", date: "340" },
        { id: 1, parameters: "HGB", date: "" },
        { id: 1, parameters: "HCT", date: "" },
        { id: 1, parameters: "MCH", date: "" },
        { id: 1, parameters: "PLT", date: "" },
      ],
    };
  },

  //Validation
  // computed: {},
  //
  methods: {
    async onShowSummaryPDF() {
      const result = await this.api.get('/report/SummaryReport', { params: { DocID: String(localStorage.getItem("VDDOCUMENT")) } })
      let header = result.data.header
      let consultant = result.data.consultant
      let documentID = result.data.documentID
      let prescriptionsSrc = result.data.prescriptions === undefined ? [] : result.data.prescriptions
      let problemListSrc = result.data.problemList === undefined ? [] : result.data.problemList
      let profileData = this.Profile[0]
    // A4 dimensions : 210 x 297
    const doc = new this.$PDF()
    doc.setProperties({
        title: 'Summary Report',
        subject: 'Report',
        author: 'admin',
        keywords: '',
        creator: 'SJGH'
    })

    let img = new Image()
    img.src = '/sjgh/logo.png'
    doc.addImage(img, 'png', 15, 15, 20, 20)
    let problemsList = []
    let drugList = []
    let investigationList = []
    if(problemListSrc !== null){
      problemListSrc.forEach((item, index, arr) => {
          problemsList.push([{
              content: `•	${item.ProblemDescription === null || item.ProblemDescription === undefined ? '' : item.ProblemDescription} - ${item.Note === null || item.Note === undefined ? '' : item.Note}`, colSpan: 6
          }])
      });
    }
    if(prescriptionsSrc !== null){
      prescriptionsSrc.forEach((item, index, arr) => {

          if (!item.Urgent) {
              let ID = item.ID;
              let code = item.ItemCode;
              let comment = item.comment === null || item.comment === undefined ? '' : item.comment;
              let description = item.Description;
              let dosage = item.Dosage;
              let duration = item.duration;
              let strenth = item.Strenth;
              if (item.ItemType.trim() === "DRUG") {

                  drugList.push([{
                      content: `${description === null ? '' : description}\n${comment === null ? '' : comment}`, colSpan: 3
                  }, `${strenth === null ? '' : strenth}`, `${dosage === null ? '' : dosage}`, `${duration === null ? '' : duration}`]);

              } else {
                  investigationList.push([{
                      content: `${description === null ? '' : (description.trim().toLowerCase().includes("other") ? '' : description)} - ${comment === null ? '' : comment}`, colSpan: 6
                  }]);

              }
          }
      });
    }
    drugList.push([{
        content: '', colSpan: 3
    }, '', '', '']);
    investigationList.push([{
        content: '', colSpan: 6
    }]);
    problemsList.push([{
        content: '', colSpan: 6
    }]);
    doc.autoTable({
        margin: { right: 0, left: 0 },
        styles: {
            fontStyle: 'normal',
            cellPadding: 0.01,
        },
        columnStyles: {
            0: { halign: 'left', cellWidth: 25 },
            1: { halign: 'center' }
        },
        theme: 'plain',
        body: [
            ['', { content: 'Endocrinology & Diabetes Centre', styles: { fontStyle: 'bold',fontSize:18}}],
            ['', 'Sri Jayewardenepura General Hospital'],
            ['', 'Thalapathpitiya, Nugegoda, Sri Lanka.'],
            ['', 'Tel: 0112 778 610 Fax: 0112 778 611 Email: endocrine@sjghsrilanka.lk Web: sjgh.health.gov.lk	']
        ]
    })
    doc.autoTable({
        margin: { right: 0, left: 0 },
        styles: {
            fontStyle: 'bold',
            fontSize: 18,
            cellPadding: 0.01,
        },
        columnStyles: {
            0: { halign: 'center' }
        },
        theme: 'plain',
        body: [
            ['Clinical Summary']
        ]
    })
    doc.autoTable({
        margin: { right: 2, left: 2 },
        styles: {
            fontStyle: 'normal',
            cellPadding: 1,
            lineWidth: 0.01,
            lineColor: [0,0,0]
        },
        columnStyles: {
            0: { halign: 'left', cellWidth: 35 },
            1: { halign: 'left' },
            2: { halign: 'left', cellWidth: 35 },
            3: { halign: 'left', cellWidth: 35 },
            4: { halign: 'left', cellWidth: 35 },
            5: { halign: 'left'},
        },
        theme: 'plain',
        body: [
            [{ content: 'Date', styles: { fontStyle: 'bold' } }, DateTime.fromISO(header.EntryDate.trim()).toFormat('dd-MM-yyyy hh:mm'), { content: 'Document No', styles: { fontStyle: 'bold' }}, { content: documentID.trim(),colSpan:3}],
            [{ content: 'Registration No ', styles: { fontStyle: 'bold' } }, header.RegistrationNo.trim(), { content: 'Age', styles: { fontStyle: 'bold' } }, profileData.Age],
            [{ content: 'Patient Name', styles: { fontStyle: 'bold' } }, { content: profileData.PatientName }, { content: 'BP', styles: { fontStyle: 'bold' } }, { content: header.BP }, { content: 'Pulse Rate', styles: { fontStyle: 'bold' } }, { content:  header.PulseRate }],
            [{ content: 'Problem List', colSpan: 6, styles: { fontStyle: 'bold' }}],
            ...problemsList,
            [{ content: '', colSpan: 6 }],
            [{ content: 'Notes', colSpan: 6, styles: { fontStyle:'bold'}}],
            [{ content: header.Comment, colSpan: 6 }],
            [{content: '', colSpan: 6 }],
            [{ content: 'Prescription', colSpan: 6, styles: { fontStyle: 'bold' } }],
            [{ content: 'Rx', colSpan: 3, styles: { fontStyle: 'bold' } }, { content: 'Dosage', styles: { fontStyle: 'bold' } }, { content: 'Frequency', styles: { fontStyle: 'bold' } }, { content: 'Duration', styles: { fontStyle: 'bold' } }],
            ...drugList,
            [{ content: '', colSpan: 6 }],
            [{ content: 'Required Investigations for Next Visit', colSpan: 6, styles: { fontStyle: 'bold' } }],
            ...investigationList,
            [{ content: 'Prepared By: ' + ' TODO ', colSpan: 3 },  { content: 'Referrals:', colSpan: 3 }],
            [{ content: (consultant.Name === null || consultant.Name === undefined ? '' : consultant.Name )+ '\n' + ( consultant.Qualifications === null || consultant.Qualifications === undefined ? '' : consultant.Qualifications )+ '\nConsultant\nSri Jayawardanapura General Hospital\nNugegoda, Sri Lanka', colSpan: 3, rowSpan: 2 }, { content: 'Next Clinic visit in: ' + (header.NextClinicIn === null ? '' : header.NextClinicIn) , colSpan: 3 }],
            [{ content: 'Annual Visit: ' + (header.AnnualVisit === null ? '' : header.AnnualVisit), colSpan: 3}],
        ]
    })
    doc.output('dataurlnewwindow');
    },

    async registrationWiseClinicSession() {
      let intial = {"label": "New","value":"--"}

      localStorage.getItem("RegiNo");
      let regi, clinic;
      regi = String(localStorage.getItem("RegiNo"));
      clinic = String(localStorage.getItem("Location"));
      const result = await this.api.get(
        `/Clinic-notes/RegistrationWiseClinicSessions?regno=${regi}&clinic=${clinic}`
      );
       this.ClinicSession = result.data.RegistrationWiseClinicSessions;
       
        
      if (this.ClinicSession != null) {
        this.ClinicSession = this.ClinicSession.map((x) => ({
          label: x.DocumentNo,
          value: x.DocumentNo,
        }));
        this.ClinicSession.push(intial)
        this.ClinicWiseSession = String(localStorage.getItem("VDDOCUMENT"))
      }else{
        this.ClinicSession= this.NewSession
      }
    },

    // Done
    async GetDocumentHeaderProblemList() {
      let DocNo;
      DocNo = String(localStorage.getItem("VDDOCUMENT"));
      //  GET DOCUMENT HEADER PROBLEM LIST
      const result = await this.api.get(
        `/Clinic-notes/GetDocumentHeaderProblemList?DocID=${DocNo}`
      );
      this.problemList =
        result.data.GetDocumentHeaderProblemList ?? [];
        for (let i = 0; i < this.problemList.length; i++) {
          this.problemList[i].Note = this.problemList[i].Note ?? ''  
        }
        let priorities = this.problemList.map(x=>x.Priority)
        this.diagnosis.ProblemPriority = Math.max( ...(priorities) ) + 1
    },
    //ResetInputField
     ResetInput(){
        header.bp='';
        header.pulse_rate= '',
        header.weight= 0;
        header.height= 0;
        header.bmi= 0;
        header.classificationValue= null;
        header.status= null;
        header.status_com= null;
        header.comment_print= null;
        header.comment_internal= null;
        header.nextClinicIn= null;
        header.nextClinicInValue= null;
        header.nextClinicInUnit= null;
        header.annualVisit= null;
     },
     async GetDocumentCheckedList() {
      let DocNo;
      DocNo = String(localStorage.getItem("VDDOCUMENT"));
      //  GET DOCUMENT HEADER PROBLEM LIST
      const result = await this.api.get(
        `/Clinic-notes/GetDocumentCheckedList?DocID=${DocNo}`
      );
      this.checkedList =
        result.data.GetDocumentCheckedList;
        this.checkedList.sort(function(a, b){
          if(a.ItemCode < b.ItemCode) { return -1; }
          if(a.ItemCode > b.ItemCode) { return 1; }
          return 0;
      })
      for (let index = 0; index < this.checkedList.length; index++) {
        let date = this.checkedList[index].Date 
        if(date !== null){
          this.checkedList[index].value = true
        }else{
          this.checkedList[index].value = false  
        }
      }
    },
     async GetDocumentPatientCheckedList() {
      let regi, clinic;
      regi = String(localStorage.getItem("RegiNo"));
      clinic = String(localStorage.getItem("Location"));
      //  GET DOCUMENT HEADER PROBLEM LIST
      const result = await this.api.get(
        `/Clinic-notes/GetDocumentPatientCheckedList?RegiNo=${regi}&Clinic=${clinic}`
      );
      this.checkedGrid =
        result.data.GetDocumentPatientCheckedList;
      if(this.checkedGrid !== null){
        let keys = Object.keys(this.checkedGrid[0])
        this.checkedGridFields = []
        for (let i = 0; i < keys.length; i++) {
          const key = keys[i];
          let item =  { filter: false, sorter: false, key: key, label: key };
          this.checkedGridFields.push(item);
        }
        for (let i = 0; i < this.checkedGrid.length; i++) {
          for (let t = 1; t < keys.length; t++) {
            let key = keys[t];
            this.checkedGrid[i][key] = this.checkedGrid[i][key] !== null ? '✔️' : '';
          }
        }
      }
    },
    async show() {
      let DocNo = String(localStorage.getItem("VDDOCUMENT"));
      const result = await this.api.get(
        `Clinic-notes/GetDocumentHeaderNotes?DocID=${DocNo}`
      );
      //  = ;
      let BodyMI = Object.assign(result.data.GetDocumentHeaderNotes);

      this.header.bp = BodyMI[0].BP;
      this.header.pulse_rate = BodyMI[0].PulseRate;
      this.header.weight = BodyMI[0].Weight;
      this.header.height = BodyMI[0].Height;
      this.header.bmi = BodyMI[0].BMI;
      this.header.classificationValue = BodyMI[0].PulseRate;
      //  Have Find below value
      //  this.header.status = BodyMI[0].PulseRate
      //  this.header.status_com = BodyMI[0].PulseRate

      this.header.comment_print = BodyMI[0].Comment;
      this.header.comment_internal = BodyMI[0].CommentInternal;
      this.header.annualVisit = BodyMI[0].AnnualVisit;
      this.header.nextClinicIn = BodyMI[0].NextClinicIn;
      this.header.nextClinicInValue = this.header.nextClinicIn.split(' ')[0]
      this.header.nextClinicInUnit = this.header.nextClinicIn.split(' ')[1]
      // this.BodyMI
    },
    // ==================================================
    async GetClinicMasterData() {
      let ClinicCode,
        DocumentID,
        DocID = String(localStorage.getItem("DocId"));
      DocumentID = String(localStorage.getItem("VDDOCUMENT"));
      // DocumentID = "0039MC12/03/22      ";
      const resultForDosage = await this.api.get(`/Clinic-notes/GetDosageList`);
      this.GetDosageList = resultForDosage.data.GetDosageList
      .map((data) => ({
       label: data.vcItemDescrition.trim(),
       value: data.cItemCode.trim(),
      }));

      const resultForUnit = await this.api.get(`/Clinic-notes/GetUnitList`);
      this.GetUnitList = resultForUnit.data.GetUnitList.map((data) => ({
        label: data.cItemCode,
        value: data.vcItemDescrition,
      }));
      //
      const resultForPrescription = await this.api.get(
        `/Clinic-notes/GetDocumentPrescription?DocID=${DocumentID}`
      );
      this.DataforPrescriptons =
        resultForPrescription.data.GetDocumentPrescription ?? [];
      for(let i = 0 ; i < this.DataforPrescriptons.length ; i++) {
        let data = this.DataforPrescriptons[i]
        this.DataforPrescriptons[i].Strenth = (data.Strenth === 'NA'|| data.Strenth === null ) ? null : data.Strenth.trim()
        this.DataforPrescriptons[i].Dosage =  (data.Dosage === 'NA' || data.Dosage === null ) ? null : data.Dosage.trim()
        this.DataforPrescriptons[i].duration =  (data.duration === 'NA'|| data.duration === null ) ? null : data.duration.trim()
      }
      this.Tablelist();
      this.ReferralTemplates();
      this.ReferralForms();
    },
    /*================================================================*/ 
    // REFERRAL TEMPLATED
     /*================================================================*/ 
    async ReferralTemplates(){
       let ClinicCode = String(localStorage.getItem("Location"));
        const resultForReferralTemp = await this.api.get(`/Clinic-notes/GetDocumentReferralLetterTemplate?clinic=${ClinicCode}`);
        var data =  resultForReferralTemp.data.GetDocumentReferralLetterTemplate;
        if(data!=null){
            this.ReferralLetterTemplate = resultForReferralTemp.data.GetDocumentReferralLetterTemplate;
        }
      
    },
   /*================================================================*/ 
   // REFERRAL FORMS
   /*================================================================*/ 
    async ReferralForms(){
      let DocumentID = String(localStorage.getItem("DDOCUMENT"));
      const resultForReferralLetter = await this.api.get(`/Clinic-notes/GetDocumentReferralLetter?DocumentID=${DocumentID}`);
      var data = resultForReferralLetter.data.GetDocumentReferralLetter
      if(data!=null){
         this.ReferralLetter = data.map((data) => ({label: data.LetterID,value: data.LetterID,}));
         return this.ReferralLetterform = data
      }

    },

    async Promise() {
      let RegiNo = localStorage.getItem("RegiNo");
      const resultGetPatientRegistionInfo = await this.api.get(
        `/reception/GetPatientRegistrationInfo?RegiNo=${RegiNo}`
      );
      this.Profile =
        resultGetPatientRegistionInfo.data.GetPatientRegistrationInfo;
    },
    //====================================================
    async GetDiseaseofImage() {
      let DocNo = this.$route.params.id;
      const resultImageOfDisease = await this.api.get(
        `Clinic-notes/GetPatientDiseaseOfImage?RegiNo=${DocNo}`
      );
      this.ImageOfDisease = resultImageOfDisease.data.GetPatientDiseaseofImage;
    },

    async GetProfile() {
      let RegiNo = localStorage.getItem("RegiNo");
      const result = await this.api.get(
        `Clinic-notes/GetPatientProfileImage?RegiNo=${RegiNo}`
      );
      let ProfileImage = result.data.GetPatientProfileImage;
      this.Photo = ProfileImage[0].Patient_Image;
    },
    //Initial Screen
    async GetDocumentInitialScreening() {
      let clinicCode = localStorage.getItem("Location");
      let registrationNo = localStorage.getItem("RegiNo");
      const result = await this.api.get(
        `Clinic-notes/GetDocumentInitialScreening?RegiNo=${registrationNo}&ClinicCode=${clinicCode}`
      );
      let initData = result.data.GetDocumentInitialScreening;
      this.initScreen.id = initData[0].ID;
      this.initScreen.presentcomplaint = initData[0].PresentingComplaint;
      this.initScreen.past_medi_or_surgicalHx = initData[0].PastMedicalHx;
      this.initScreen.obstretic_or_menstrualH = initData[0].ObstreticHx;
      this.initScreen.medication_summary = initData[0].MedicationSummary;
      this.initScreen.familtyHistory = initData[0].FamilyHistory;
      this.initScreen.allergicHx = initData[0].AllegicHx;
      this.initScreen.socialHx = initData[0].SocialHx;
      this.initScreen.generalExamination = initData[0].GeneralExamination;
      this.initScreen.smorking = initData[0].Smoking;
      this.initScreen.alcohol = initData[0].Alcohole;
      this.initScreen.fGS = initData[0].FGS;
      this.initScreen.fGS_Image = initData[0].FGSImage;
      this.initScreen.Thyroid_Examination = initData[0].ThyroidExam;
      this.initScreen.eyes = initData[0].Eyes;
      this.initScreen.rS = initData[0].RS;
      this.initScreen.aBD = initData[0].ABD;
      this.initScreen.cNS = initData[0].CNS;
      this.initScreen.cVS = initData[0].CVS;
    },

    reset_comment() {
      this.comment.comment_print = null;
      this.comment.comment_internal = null;
    },
    reset_health() {
      (this.header.bp = null),
        (this.header.pulse_rate = null),
        (this.header.weight = null),
        (this.header.height = null),
        (this.header.bmi = null),
        (this.header.status = null),
        (this.header.status_com = null),
        (this.header.class = null);
    },
    reset_referral() {
      (this.ReferralForm.referral_id = null),
        (this.ReferralForm.rdoctor = null),
        (this.ReferralForm.rnotes = null);
    },

    addSpecialNote(diagnosis,problemList) {
      if(diagnosis.index !== null){
      problemList[diagnosis.index].ProblemCode = diagnosis.ProblemDescriptionValue.ID;
      problemList[diagnosis.index].ProblemDescription =  diagnosis.ProblemDescriptionValue.Description ?? '';
      problemList[diagnosis.index].Note = diagnosis.ProblemNote ?? '';
      problemList[diagnosis.index].Status = diagnosis.ProblemStatus ? 'Active' : 'Deactive';
      problemList[diagnosis.index].Priority = diagnosis.ProblemPriority;
      }else{
      let item = {};
      item.index =  problemList.length ;
      item.ProblemCode = diagnosis.ProblemDescriptionValue.ID;
      item.ProblemDescription =  diagnosis.ProblemDescriptionValue.Description ?? '';
      item.Note = diagnosis.ProblemNote ?? '';
      item.Status = diagnosis.ProblemStatus ? 'Active' : 'Deactive';
      item.Priority = diagnosis.ProblemPriority;
      problemList.push(item)
      }
      diagnosis.ProblemNote = ''
      diagnosis.ProblemDescriptionValue = null
      diagnosis.ProblemStatus = null
      diagnosis.ProblemPriority = Math.max(...problemList.map(o => o.Priority), 0) + 1;
      diagnosis.index = null;
    },
    editSpecialNote(item,problemList) {
      this.diagnosis.index = problemList.indexOf(item)
      this.diagnosis.ProblemNote = item.Note ?? '';
      this.diagnosis.ProblemStatus = item.Status.trim() === 'Active' ? true : false;
      this.diagnosis.ProblemPriority = item.Priority;
      this.diagnosis.ProblemDescriptionValue = this.diagnosticList.find( x => x.value.ID === Number(item.ProblemCode)).value;
    },
    deleteSpecialNote(item,problemList) {
      this.problemList = problemList.filter(a => !(a === item));
      this.deleteDiagnosisItem = null
      this.deleteDiagnosisModal = false
      this.GetDocumentHeaderProblemList()
    },
    ad() {
      alert("hello");
    },
    deleteinfo() {
      alert("Delete function is unable fire");
    },
    GetUsertypes() {
      let user = localStorage.getItem("Usertype");
      this.Usertype = user;
    },

    imageload() {
      // alert("image")
      this.warningModal = true;
    },
    //*==============================================================================*/
    //                         SAVE AS REFERRAL FORM                                 //
    /*===============================================================================*/
    async SaveasReferralform(){
      let document = String(localStorage.getItem("VDDOCUMENT"))
      let form ={}
      form.DocumentID = document
      form.LetterID = this.PatientReferralbyID.length+1
      form.DoctorID= localStorage.getItem("UserId");
      form.Letter = this.content
      // form.ClinicID = localStorage.getItem("Location");
      this.PatientReferralbyID.push(form)
      console.log("REFERRAL",this.ReferralForm.Letter)
    },
    

  /*===============================================================================*/
  //                          SAVE AS TEMPLATE                                    //
  /*===============================================================================*/
    async SaveasTemplate(){ 
       let form ={}
         form.TemplateName = this.ReferralTemp
         form.DoctorID = localStorage.getItem("UserId")
         form.Letter = this.content
         form.ClinicID = localStorage.getItem("Location")
        this.ReferralLetterTemplateArray.push(form)
       setTimeout(()=>(this.ReferraltemplateName=false),1500);
    },

   /* ==== Call referral template modal ==== */
    async CallRefNameTempModel(){
       this.ReferraltemplateName=true
    },
    //==========

    loadRefTemplate(item) {
      this.ReferralForm.Letter = item.Letter;
      setTimeout((this.loadtemplate = false), 1500);
    },
    close() {
      // alert()
      this.warningModal = false;
    },
    uploadImage() {
      // alert()
      this.imageModal = true;
    },
    openUpload() {
      document.getElementById("file").click();
    },
    UploadPreview(e) {
      var reader,
        files = e.target.files;
      reader = new FileReader();
      reader.onload = (e) => {
        let img = e.target.result;

        this.imagePreviewFordisease = img;
        // let cropper = this.imagePreview
        // let  crop = new Cropper(img,{
        //   zoomable:false,
        //   scalable:false,
        //   aspectRatio:1
        // })
      };
      reader.readAsDataURL(files[0]);
    },
    calcBMI() {
      this.header.bmi = this.header.weight / ((this.header.height/100) * (this.header.height/100))
      this.header.bmi = Math.round(this.header.bmi * 100) / 100
    },
    check(item) {
      const val = Boolean(this.usersData[item.id]._selected);
      this.$set(this.usersData[item.id], "_selected", !val);
    },

    closeIm() {
      this.imageModal = false;
    },

    loadTemplate() {
      this.loadtemplate = true;
    },

    GetPatientReferral(item) {
    },
    modaltest() {
      this.mytext = true;
      alert();
    },

    /*================================================*/
    //                   DOC Number SELECT            //
    /*================================================*/ 
    changeValue() {
      let DocumentNo = this.ClinicWiseSession;
      localStorage.setItem("VDDOCUMENT", DocumentNo);
      this.iternalReload()
      this.show()
      if(DocumentNo === "New"){
        this.show()
        this.iternalReload()
      }
      else{
        let route_params = this.$route.params.id;
          let route = this.$router.resolve({path: `/doctor-queue-or-portal_/${route_params}/Notes`,});
          window.open(route.href, '_blank');
      }
    },

   
    async iternalReload(){
      this.GetDocumentHeaderProblemList()
      this.GetClinicMasterData()
    },
    
    async onDrugSelect() {
      if ("Drugs" === this.Prescript.drugsSearch) {
        this.DrugsPropertyAutocomplete()
      }
    },
    async onDrugSearch(e) {
      // let val = e.split(" ").slice(1).join(" ") ?? e;
       let val = this.Prescript.InputSearch
       let Stext= val.length;
      //  console.log(val,"SEAkkoRCH",Stext)
       let i=3
       
      if ("Drugs" === this.Prescript.drugsSearch) {
        if(i<Stext){
        let text = val
        const result = await this.api.get(`/Clinic-notes/GetFilterDruglist`,{params: { SearchText: text}});
        const filterd = result.data.GetFilterDrugsList;
        let mo = filterd.map((data) => {
          let item = data
          item.label = data.ATCClessification.trim()
          item.value = data.cItemCode.trim() + ' ' + data.ATCClessification.trim()
          item.itemCode = data.cItemCode.trim()
          return item
        });
        return this.InputSearchSelect = mo;
       }
      } 
      else if ("Lab" === this.Prescript.drugsSearch) {
        let text = val
        const result = await this.api.get(`/Clinic-notes/GetFilterLabInvestlist`,{params:{SearchText:text}});
        const filterd = result.data.GetFilterLablist;
        let mo = filterd.map((data) => {
          let item = data
          item.label = data.vcItemDescrition.trim()
          item.value = data.cItemCode.trim() + ' ' + data.vcItemDescrition.trim()
          item.itemCode = data.cItemCode.trim()
          return item
        });
        this.InputSearchSelect = mo;
      } 
      else if ("Other" === this.Prescript.drugsSearch) {
        let text = val
        const result = await this.api.get(`/Clinic-notes/GetFilterOtherlist`,{params:{SearchText:text}});
        const filterd = result.data.GetFilterOtherInvestigationlist;
        let mo = filterd.map((data) => {
          let item = data
          item.label = data.vcItemDescrition.trim()
          item.value = data.cItemCode.trim() + ' ' + data.vcItemDescrition.trim()
          item.itemCode = data.cItemCode.trim()
          return item
        });
        this.InputSearchSelect = mo;
      }
    },
    DisabledForPrescriptionOtherField(){
      if ("Lab" === this.Prescript.drugsSearch){
         document.getElementById("Strenth").disabled = true;
         document.getElementById("Unit").disabled = true;
         document.getElementById("Frequncy").disabled = true;
         document.getElementById("Duration").disabled = true;
         document.getElementById("Urgent").disabled = false;

      }
      else if ("Other" === this.Prescript.drugsSearch){
        document.getElementById("Strenth").disabled = true;
        document.getElementById("Unit").disabled = true;
        document.getElementById("Frequncy").disabled = true;
        document.getElementById("Duration").disabled = true;
         document.getElementById("Urgent").disabled = false;
               
      }

       else if ("Drugs" === this.Prescript.drugsSearch){
        document.getElementById("Strenth").disabled = false;
        document.getElementById("Unit").disabled = false;
        document.getElementById("Frequncy").disabled = false;
        document.getElementById("Duration").disabled = false;  
        document.getElementById("Urgent").disabled = false;
      }
    },
    async DrugsPropertyAutocomplete(){
        const INPUTTEXT = this.Prescript.InputSearch.split(" ").splice(1).join(" ") ?? this.Prescript.InputSearch
        const result = await this.api.get(`/Clinic-notes/GetDrugAutoComplete`,{params: { GenericName: INPUTTEXT }})
        const drugautoComplete = result.data.DrugAutoCompleteI
        this.Prescript.selectedDrug = drugautoComplete[0]
        let unit = this.Prescript.selectedDrug.Unit
        this.Prescript.strenth = (this.Prescript.selectedDrug.Strength ?? "").trim()
        this.Prescript.selUnit = (unit ?? "").trim()
    },
    AddForPrescription(){
      let type = this.Prescript.drugsSearch.trim()
      switch (type) {
        case "Drugs":
          type =  "DRUG"   
          break;
        case "Lab":
          type = "LAB"
          break;
        case "Other":
          type = "OTHER"
          break;
        default:
          break;
      }
      if(this.Prescript.index === null){
        let item = {};
        item.documentID = localStorage.getItem("VDDOCUMENT");
        item.ID  = 0;
        item.ItemCode=this.Prescript.selectedDrug == null ? (this.Prescript.InputSearch.split(' ')[0]  ?? this.Prescript.InputSearch) : this.Prescript.selectedDrug.GenericId;
        item.comment = this.Prescript.comment;
        item.ItemType = type.trim();
        item.Description = this.Prescript.InputSearch;
        if(item.ItemType.trim() === "DRUG"){
        item.Dosage = this.Prescript.selFreque ;
        item.duration = this.Prescript.week;
        item.Strenth = this.Prescript.strenth + " " + this.Prescript.selUnit;
        }else{
        item.Dosage = "NA"
        item.duration = "NA"
        item.Strenth = "NA"  
        }
        item.Urgent = this.Prescript.urgentCheck  ?? false;
        item.Sequence = (this.DataforPrescriptons === null ) ? 1 : (this.DataforPrescriptons.length + 1) ;
        this.DataforPrescriptons.push(item);
      }else{
        let index = this.Prescript.index
        this.DataforPrescriptons[index].CompanyCode = this.Prescript.drugsSearch
        this.DataforPrescriptons[index].DocumentID = String(localStorage.getItem("VDDOCUMENT"))
        this.DataforPrescriptons[index].ItemCode = this.Prescript.selectedDrug == null ? (this.Prescript.InputSearch.split(' ')[0]  ?? this.Prescript.InputSearch) : this.Prescript.selectedDrug.GenericId
        this.DataforPrescriptons[index].ItemType = type.trim()
        this.DataforPrescriptons[index].Description = this.Prescript.InputSearch
        this.DataforPrescriptons[index].comment = this.Prescript.comment
        if(this.DataforPrescriptons[index].ItemType.trim() === "DRUG"){
        this.DataforPrescriptons[index].Dosage = this.Prescript.selFreque ;
        this.DataforPrescriptons[index].duration = this.Prescript.week;
        this.DataforPrescriptons[index].Strenth = this.Prescript.strenth + " " + this.Prescript.selUnit;
        }else{
        this.DataforPrescriptons[index].Dosage = "NA"
        this.DataforPrescriptons[index].duration = "NA"
        this.DataforPrescriptons[index].Strenth = "NA"  
        }
        this.DataforPrescriptons[index].Urgent = this.Prescript.urgentCheck  ?? false
        this.DataforPrescriptons[index].Sequence = (this.DataforPrescriptons === null ) ? 1 : (this.DataforPrescriptons.length + 1) 
      }
      this.InputSearchSelect = null
      this.Prescript.index = null
      this.Prescript.selectedDrug = null
      this.Prescript.selFreque  = null
      this.Prescript.InputSearch  = null
      this.Prescript.comment  = null
      this.Prescript.selUnit  = null
      this.Prescript.strenth = null 
      this.Prescript.urgentCheck   = false
      this.Prescript.week   = null
      this.DisabledForPrescriptionOtherField()
      this.Tablelist()
    },
    
    editPrescription(vitem){
      let item = this.DataforPrescriptons[vitem.idx]
      this.Prescript.index = this.DataforPrescriptons.indexOf(item)
      this.Prescript.InputSearch = item.Description
      this.Prescript.id = item.ID
      let type = item.ItemType.trim()
      switch (type) {
        case "DRUG":
          this.Prescript.drugsSearch =  "Drugs"   
          break;
        case "LAB":
          this.Prescript.drugsSearch = "Lab"
          break;
        case "OTHER":
          this.Prescript.drugsSearch = "Other"
          break;
        default:
          break;
      }
      
      this.Prescript.comment = item.comment 
      this.Prescript.urgentCheck = item.Urgent
      if(type === "DRUG"){
        this.Prescript.selFreque = item.Dosage.trim() ?? ""
        this.Prescript.week = item.duration
        this.Prescript.strenth = item.Strenth.split(" ")[0]
        this.Prescript.selUnit = item.Strenth.split(" ")[1]
      }else{
        this.Prescript.selFreque = null
        this.Prescript.week = null
        this.Prescript.strenth = null
        this.Prescript.selUnit = null
      }
      this.DisabledForPrescriptionOtherField()
    },

    async deletePrescription(vitem){
      let item = this.DataforPrescriptons[vitem.idx]
      this.problemList = this.DataforPrescriptons.filter(a => !(a === item));
      this.deletePrescriptionItem = null
      this.deletePrescriptionModal = false
      this.DisabledForPrescriptionOtherField()
      this.Tablelist()
    },

    async Tablelist(){
      /* PRESCRIPTIONS */ 
      this.DataforPrescriptonslist = this.DataforPrescriptons.filter((data) => (data.Urgent !==true)).map((data) => ({Dosage: (data.Strenth ?? 'NA') +" | "+ (data.Dosage ?? 'NA') +" | "+ (data.duration ?? 'NA'),Comment: data.comment,Description: data.Description, ItemType: data.ItemType, ID:data.ID, idx: this.DataforPrescriptons.indexOf(data)}));

      /* URGENT PRESCRIPTIONS */ 
      this.UrgentPrescriptionDatalist = this.DataforPrescriptons.filter((data) => (data.Urgent ===true)).map((data) => ({Dosage: (data.Strenth ?? 'NA') +" | "+ (data.Dosage ?? 'NA') +" | "+ (data.duration ?? 'NA'),Comment: data.comment,Description: data.Description, ItemType: data.ItemType, ID:data.ID, idx: this.DataforPrescriptons.indexOf(data)}));
      
    },
 
   
    // Diagnosis

    async GetDiagnosticList() {
      const result = await this.api.get(
        `Clinic-notes/GetDiagnosticList`
      );
      this.diagnosticList = result.data.GetDiagnosticList.map((data) => ({
        label: data.Description,
        value: data,
      }));
    },
    async onSaveClick() {
      await this.SaveClinicHeader(); 
      if (this.tabIndex == 5) {
        await this.SaveClinicInitScreening();
      }
      if (this.tabIndex == 0) {
        await this.SaveClinicSpecialNotes();
      }
      if (this.tabIndex == 6) {
        await this.SaveClinicCheckedList();
      }
      if(this.tabIndex == 2 ) {
        await this.SaveClinicPrescription();
      }
      if(this.tabIndex == 4){
         await this.SaveClinicReferral();
         await this.SaveClinicReferralTemplate();
      }
    },
    async SaveClinicHeader() {
      this.header.nextClinicIn = `${this.header.nextClinicInValue} ${this.header.nextClinicInUnit}`;
      let form = {};
      form.registrationNo = String(localStorage.getItem("RegiNo"));
      form.documentID = String(localStorage.getItem("VDDOCUMENT"));
      form.userID = String(localStorage.getItem("UserId"));
      form.BP = (this.header.bp ?? "").trim();
      form.pulseRate = this.header.pulse_rate ?? 0;
      form.weight = this.header.weight ?? 0;
      form.height = this.header.height ?? 0;
      form.BMI = this.header.bmi ?? 0;
      form.comment = (this.header.comment_print ?? "").trim();
      form.commentInternal = (this.header.comment_internal ?? "").trim();
      form.nextClinicIn = (this.header.nextClinicIn ?? "").trim();
      form.annualVisit = (this.header.annualVisit ?? "").trim();
      this.api
        .post("Clinic-notes/SaveClinicHeader", form)
        .then((result) => {
          if (result.data === "success") {
            this.alertboxSuccess = true;
            setTimeout(() => (this.alertboxSuccess = false), 3000);
            this.show();
          } else {
            this.alertboxFailed = true;
            setTimeout(() => (this.alertboxFailed = false), 3000);
          }
        })
        .catch((error) => {
          this.err = error;
          this.alertboxFailed = true;
          setTimeout(() => (this.alertboxFailed = false), 3000);
        });
    },
    async SaveClinicInitScreening() {
      let form = {};
      form.id = this.initScreen.id;
      form.presentingComplaint = this.initScreen.presentcomplaint ?? '';
      form.pastMedicalHx = this.initScreen.past_medi_or_surgicalHx ?? '';
      form.obstreticHx = this.initScreen.obstretic_or_menstrualH ?? '';
      form.MedicationSummary = this.initScreen.medication_summary ?? '';
      form.familyHistory = "";
      form.allegicHx = this.initScreen.allergicHx?? '';
      form.socialHx = this.initScreen.socialHx?? '';
      form.occupation = "";
      form.smoking = this.initScreen.smorking ?? '';
      form.alcohol = this.initScreen.alcohol ?? '';
      form.generalExamination = this.initScreen.generalExamination ?? '';
      form.FGS = this.initScreen.fGS ?? '';
      form.FGSImage = this.initScreen.fGS_Image ?? '';
      form.ThyroidExam = this.initScreen.Thyroid_Examination ?? '';
      form.Eyes = this.initScreen.eyes ?? '';
      form.RS = this.initScreen.rs ?? '';
      form.ABD = this.initScreen.aBD ?? '';
      form.CNS = this.initScreen.cNS ?? '';
      form.CVS = this.initScreen.cVS ?? '';
      form.clinicType = String(localStorage.getItem("Location"));
      form.registrationID = String(localStorage.getItem("RegiNo"));
      form.userID = String(localStorage.getItem("UserId"));
      this.api
        .post("Clinic-notes/SaveClinicInitScreening", form)
        .then((result) => {
          if (result.data === "success") {
            this.alertboxSuccess = true;
            setTimeout(() => (this.alertboxSuccess = false), 3000);
            this.show();
          } else {
            this.alertboxFailed = true;
            setTimeout(() => (this.alertboxFailed = false), 3000);
          }
        })
        .catch((error) => {
          this.err = error;
          this.alertboxFailed = true;
          setTimeout(() => (this.alertboxFailed = false), 3000);
        });
    },
    onCheckListSave() {
      this.SaveClinicCheckedList()
    },
    async SaveClinicCheckedList() {
      let form = {};
      form.documentID = String(localStorage.getItem("VDDOCUMENT"));
      let items = this.checkedList.filter(x=>x.value === true)
      form.itemIDs = items.map( x => ({
        itemID: x.ItemCode.trim()
      }));
      this.api
        .post("Clinic-notes/SaveClinicCheckList", form)
        .then((result) => {
          if (result.data === "success") {
            this.alertboxSuccess = true;
            setTimeout(() => (this.alertboxSuccess = false), 3000);
            this.show();
          } else {
            this.alertboxFailed = true;
            setTimeout(() => (this.alertboxFailed = false), 3000);
          }
        })
        .catch((error) => {
          this.err = error;
          this.alertboxFailed = true;
          setTimeout(() => (this.alertboxFailed = false), 3000);
        });
    },
    async SaveClinicPrescription() {
      let form = {};
      form.documentID = String(localStorage.getItem("VDDOCUMENT"));
      form.prescriptions = this.DataforPrescriptons;
      for(let i = 0; i < form.prescriptions.length ; i++){
        form.prescriptions[i].Sequence = i
      }
      this.api
        .post("Clinic-notes/SaveClinicPrescription", form)
        .then((result) => {
          if (result.data === "success") {
            this.alertboxSuccess = true;
            setTimeout(() => (this.alertboxSuccess = false), 3000);
            this.show();
          } else {
            this.alertboxFailed = true;
            setTimeout(() => (this.alertboxFailed = false), 3000);
          }
        })
        .catch((error) => {
          this.err = error;
          this.alertboxFailed = true;
          setTimeout(() => (this.alertboxFailed = false), 3000);
        });
    },
    async SaveClinicSpecialNotes() {
      let form = {};
      form.registrationID = String(localStorage.getItem("RegiNo"));
      form.documentID = String(localStorage.getItem("VDDOCUMENT"));
      form.specialnotes = this.problemList;
      this.api
        .post("Clinic-notes/SaveClinicSpecialNotes", form)
        .then((result) => {
          if (result.data === "success") {
            this.alertboxSuccess = true;
            setTimeout(() => (this.alertboxSuccess = false), 3000);
            this.show();
          } else {
            this.alertboxFailed = true;
            setTimeout(() => (this.alertboxFailed = false), 3000);
          }
        })
        .catch((error) => {
          this.err = error;
          this.alertboxFailed = true;
          setTimeout(() => (this.alertboxFailed = false), 3000);
        });
    },
    /*====================================================================*/ 
//             FINAL SAVE FOR REFERRAL FORM
/*====================================================================*/ 

    async SaveClinicReferral() {
      let form = {};
      form.documentID = String(localStorage.getItem("VDDOCUMENT"));
      form.Referral = this.PatientReferralbyID
      this.api
        .post("Clinic-notes/SaveClinicReferral", form)
        .then((result) => {
          if (result.data === "success") {
            this.alertboxSuccess = true;
            setTimeout(() => (this.alertboxSuccess = false), 3000);
            this.show();
          } else {
            this.alertboxFailed = true;
            setTimeout(() => (this.alertboxFailed = false), 3000);
          }
        })
        .catch((error) => {
          this.err = error;
          this.alertboxFailed = true;
          setTimeout(() => (this.alertboxFailed = false), 3000);
        });
    },

    /*====================================================================*/ 
    //                  FINAL SAVE FOR REFERRAL TEMPLATE 
    /*====================================================================*/ 
    async SaveClinicReferralTemplate() {
      let form = {};
      form.referraltemplate=this.ReferralLetterTemplateArray
      this.api
        .post("Clinic-notes/SaveClinicReferralTemplate", form)
        .then((result) => {
          if (result.data === "success") {
            this.alertboxSuccess = true;
            setTimeout(() => (this.alertboxSuccess = false), 3000);
            this.show();
          } else {
            this.alertboxFailed = true;
            setTimeout(() => (this.alertboxFailed = false), 3000);
          }
        })
        .catch((error) => {
          this.err = error;
          this.alertboxFailed = true;
          setTimeout(() => (this.alertboxFailed = false), 3000);
        });
    },

    
    
    async SaveClinicImageDisease() {
      let form = {};
      form.doctorNoteID = String(localStorage.getItem("VDDOCUMENT"));
      form.nameOfDisease = this.specialNotes;
      form.diseaseofNote = this.specialNotes;
      form.nameofImage = this.specialNotes;
      form.userID = String(localStorage.getItem("UserId"));
      this.api
        .post("Clinic-notes/SaveClinicImageDisease", form)
        .then((result) => {
          if (result.data === "success") {
            this.alertboxSuccess = true;
            setTimeout(() => (this.alertboxSuccess = false), 3000);
            this.show();
          } else {
            this.alertboxFailed = true;
            setTimeout(() => (this.alertboxFailed = false), 3000);
          }
        })
        .catch((error) => {
          this.err = error;
          this.alertboxFailed = true;
          setTimeout(() => (this.alertboxFailed = false), 3000);
        });
    },
  },
  created() {
    this.registrationWiseClinicSession();
    this.show();
    this.GetDiagnosticList();
    this.GetDocumentHeaderProblemList();
    this.GetDocumentCheckedList();
    this.GetDocumentPatientCheckedList();
    this.GetClinicMasterData();
    this.GetDiseaseofImage();
    this.GetProfile();
    this.Promise();
    this.GetDocumentInitialScreening();
  },
};
</script>


<style></style>
