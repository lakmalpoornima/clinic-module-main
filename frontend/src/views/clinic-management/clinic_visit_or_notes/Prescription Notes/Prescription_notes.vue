<template v-for="(color, key) in ['success', 'primary', 'info', 'warning', 'danger']">
  <div>
    <CModal title="Alert" color="danger" :show.sync="alertModal">
      {{ AlertMessage }}
    </CModal>

    <CModal
      title="Insert External Lab Data"
      color="secondary"
      :show.sync="insertExternalLabModal"
      size="xl"
      class="align-self-center max-w-min"
    >
      <CRow >
        <CCol md="auto">
          Element
          <CInput type="text" v-bind:key="key" v-for="(value, key) in LabInvestInputData"  v-model="LabInvestInputData[key].label" readonly />
        </CCol>
        <CCol md="2">
          Result
          <CInput type="text" v-bind:key="key" v-for="(value, key) in LabInvestInputData" v-model="LabInvestInputData[key].result" />
        </CCol>
        <CCol md="2">
          Unit
          <CInput type="text" v-bind:key="key" v-for="(value, key) in LabInvestInputData" v-model="LabInvestInputData[key].unit" />
        </CCol>
        <CCol md="auto">
          Date
          <CInput type="date" v-bind:key="key" v-for="(value, key) in LabInvestInputData" v-model="LabInvestInputData[key].date" />
        </CCol>
      </CRow>
     <div slot="footer">
      <CRow>
        <CCol class="p-1">
          <CButton
            class="col align-self-center"
            color="secondary"
            @click="insertExternalLabModal = false"
            >Cancel</CButton
          >
        </CCol>
         <CCol class="p-1">
          <CButton
            class="col align-self-center"
            color="primary"
            @click="
              () => {
                this.SaveClinicExternalLabData();
                this.insertExternalLabModal = false;
              }
            "
            >Submit</CButton
          >
        </CCol>
      </CRow>
      </div>
    </CModal>
    
     <CModal
      title="Insert New Problem"
      color="secondary"
      :show.sync="insertNewProblemModal"
      size="xl"
      class="align-self-center max-w-min"
    >
      <CRow >
        <CCol md="8">
          Problem Description
          <CInput type="text" v-model="NewProblem.description" />
        </CCol>
      </CRow>
     <div slot="footer">
      <CRow>
        <CCol class="p-1">
          <CButton
            class="col align-self-center"
            color="secondary"
            @click="insertNewProblemModal = false"
            >Cancel</CButton
          >
        </CCol>
         <CCol class="p-1">
          <CButton
            class="col align-self-center"
            color="primary"
            @click="
              () => {
                this.AddNewProblem();
                this.insertNewProblemModal = false;
              }
            "
            >Submit</CButton
          >
        </CCol>
      </CRow>
      </div>
    </CModal>
    
    <!-- This is for referral Template Name  -->
    <CModal
      title="Referral Template"
      color="warning"
      :show.sync="ReferraltemplateName"
      class="align-self-center"
    >
      <CRow>
        <CCol>
          <CInput type="text" label="Template Name" v-model="ReferralTemp" />
        </CCol>
      </CRow>
      <div slot="footer">
        <CRow>
          <CCol class="p-1">
            <CButton
              class="col align-self-center"
              color="primary"
              @click="SaveasTemplate"
              >Save</CButton
            >
          </CCol>
          <CCol class="p-1">
            <CButton
              class="col align-self-center"
              color="secondary"
              @click="deleteDiagnosisModal = false"
              >Cancel</CButton
            >
          </CCol>
        </CRow>
      </div>
    </CModal>
    <!--  -->

    <CToaster position="top-center">
      <CToast
        :show="alertboxSuccess"
        color="primary"
        title="CToast static component"
      >
        Successful Submit
      </CToast>
      <CToast
        :show="alertboxFailed"
        color="info"
        title="CToast static component"
      >
        Failed Submit - {{ err }}
      </CToast>
    </CToaster>
    <CModal
      title="Delete Diagnosis"
      color="warning"
      :show.sync="deleteDiagnosisModal"
      class="align-self-center"
    >
      Confirm to delete Dignosis item ?
      <div slot="footer">
        <CRow>
          <CCol class="p-1">
            <CButton
              class="col align-self-center"
              color="primary"
              @click="deleteSpecialNote(deleteDiagnosisItem, problemList)"
              >Confirm</CButton
            >
          </CCol>
          <CCol class="p-1">
            <CButton
              class="col align-self-center"
              color="secondary"
              @click="deleteDiagnosisModal = false"
              >Cancel</CButton
            >
          </CCol>
        </CRow>
      </div>
    </CModal>
    <CModal
      title="Delete Prescription"
      color="warning"
      :show.sync="deletePrescriptionModal"
      class="align-self-center"
    >
      Confirm to delete Prescription item ?
      <div slot="footer">
        <CRow>
          <CCol class="p-1">
            <CButton
              class="col align-self-center"
              color="primary"
              @click="deletePrescription(deletePrescriptionItem)"
              >Confirm</CButton
            >
          </CCol>
          <CCol class="p-1">
            <CButton
              class="col align-self-center"
              color="secondary"
              @click="deletePrescriptionModal = false"
              >Cancel</CButton
            >
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
                  <div v-for="data in ImageOfDiseaseView" :key="data">
                    <img
                      v-bind:src="data.cNameofImage"
                      class="card-img-top rounded h-100"
                      alt="..."
                    />
                    <!-- {{data.cNameofdisease}} -->
                    <!-- {{data.cDiseaseofNote}} -->
                  </div>
                </div>

                <!--   -->
                <div class="col-sm-7">
                  <div v-for="data in ImageOfDiseaseView" :key="data">
                    <div class="card-body" style="margin-left: -25px">
                      <CLable color="info" class="card-title font-bolder">
                        <b>{{ data.cNameofdisease }}</b>
                      </CLable>
                      <p class="card-text">{{ data.cDiseaseofNote }}</p>
                      <div class="d-flex">
                        <p class="font-italic text-info text-wrap" color="info">
                          Examine by : {{ data.UserID }}
                        </p>
                        <p
                          class="font-italic text-yahoo text-wrap"
                          style="
                            font-size: 12px;
                            margin-top: 3px;
                            margin-left: 5px;
                          "
                        >
                          {{ new Date(data.CreateDate).toLocaleDateString() }}
                        </p>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
              <!--  -->
              <!-- </div> -->
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
                  <img
                    v-if="imagePreviewFordisease !== null"
                    id="output"
                    v-bind:src="imagePreviewFordisease"
                    class="border border-danger rounded"
                    v-on:click="openUpload"
                    style="width: 150px; height: 150px"
                  />
                  <img
                    v-else
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
                      <CLable>Disease of name</CLable>
                      <CInput type="text" v-model="disease.name" />
                    </CCol>
                  </CRow>
                  <CRow style="margin-top: -19px">
                    <CCol class="" align="left">
                      <CLable>Note </CLable>
                      <CTextarea type="text" v-model="disease.discriptions" />
                    </CCol>
                  </CRow>
                  <CRow style="margin-top: -15px">
                    <CCol
                      class="d-flex"
                      align="center"
                      style="margin-left: 10px"
                    >
                      <CButton
                        @click="diseaseformReset"
                        type="reset"
                        style="width: 8rem; margin-left: 10px"
                        color="facebook"
                        >clear</CButton
                      >
                      <CButton
                        @click="saveDiseaseofImage"
                        style="width: 8rem; margin-left: 10px"
                        color="yahoo"
                        >Save</CButton
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
                    <img
                      v-if="Photo !== null"
                      class="text-center m-0 shadow rounded"
                      v-bind:src="Photo"
                      width="85"
                      alt=""
                    />
                    <img
                      v-else
                      class="text-center m-0 shadow rounded"
                      :src="imagePreview"
                      width="75"
                      :alt="imagePreview"
                      style="height: 90px"
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
                            <span class="text-primary">{{
                              data.PatientName
                            }}</span>
                          </CLable></CCol
                        >
                      </CRow>
                      <CRow
                        style="margin-top: 1px; margin-left: -4px"
                        class="rounded"
                      >
                        <CCol
                          ><CLable
                            >Registration No<span
                              class="ml-2"
                              style="margin-right: 4px"
                            >
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
                        ><span class="text-primary">{{
                          data.Gender
                        }}</span></CLable
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
                    @change="changeValue()"
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

        <CCol md="9">
          <CCard style="margin-left: -28px; height: 100%">
            <CRow class="flex-row" style="">
              <CCol col="11" class="flex-row">
                <CTabs variant="tabs" :activeTab.sync="tabIndex">
                  <CTab title="Diagnosis">
                    <CTabContent
                      style="margin-top: 50px; padding: 0px margin-left:-30px;"
                    >
                      <CTabPane>
                        <CRow style="padding: 5px">
                          <CCol class="d-flex">
                          <CButton v-if="ShowProbAddOpt"
                              size="md"
                              color="info"
                              @click="insertNewProblemModal = true"
                            >  Add New Problem </CButton>
                          </CCol>
                        </CRow>
                        <CRow style="padding: 5px">
                          <CCol class="d-flex">
                            <CInput
                              list="diagnosis"
                              placeholder="Search DiagnosticList"
                              v-model="diagnosis.ProblemDescriptionLabel"
                              @input="GetDiagnosticList"
                              @change="onDiagnosisSelect"
                            />

                            <datalist id="diagnosis" loc>
                              <option
                                v-for="diagnostic in diagnosticList"
                                :key="
                                  diagnostic.value.ID +
                                  ' ' +
                                  diagnostic.value.Description
                                "
                                :value="diagnostic.value.Description"
                              >
                                {{
                                  diagnostic.value.ID +
                                  " " +
                                  diagnostic.value.Description
                                }}
                              </option>
                            </datalist>
                            <!--<CSelect
                                placeholder=""
                                :value.sync="diagnosis.ProblemDescriptionValue"
                                :options="diagnosticList"
                              >
                              </CSelect> -->
                            <CInput
                              v-model="diagnosis.ProblemNote"
                              placeholder="Note"
                              style="width: 40%"
                            />
                            <CInput
                              type="number"
                              v-model="diagnosis.ProblemPriority"
                              min="1"
                              style="width: 20%"
                            />
                            <CInputCheckbox
                              :checked.sync="diagnosis.ProblemStatus"
                              true-value="Active"
                              false-value="Deactive"
                            />
                            <CButton
                              size="md"
                              color="info"
                              @click="addSpecialNote(diagnosis, problemList)"
                              style="width: 8rem; height: 2rem"
                            >
                              {{
                                diagnosis.index === null ? "Add" : "Update"
                              }}</CButton
                            >
                          </CCol>
                        </CRow>

                        <CDataTable
                          style="
                            padding: 5px;
                            text-align: left;
                            margin-top: -14px;
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
                              <CButtonGroup align="right">
                                <div style="" size="sm" class="mt-6 d-flex">
                                  <i @click="editSpecialNote(item, problemList)"
                                    ><CIcon
                                      name="cil-pencil"
                                      size="sm"
                                      class="ml-2 mb-2 text-primary"
                                  /></i>

                                  <i
                                    @click="
                                      deleteDiagnosisModal = true;
                                      deleteDiagnosisItem = item;
                                    "
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
                              <CButtonGroup align="right">
                                <div style="mar" size="sm" @click="show(item)">
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
                      </CTabPane>
                    </CTabContent>
                  </CTab>

                  <CTab title="Lab Investigation">
                    <CTabContent
                      style="margin-top: 50px; padding: 0px"
                      align="center"
                    >
                      <CTabPane>
                        <CCol style="padding: 5px" class="d-flex" align="right">
                          <CButton
                            color="behance"
                            @click="
                              () => {
                                this.GetLabInvestigationData(this);
                              }
                            "
                          >
                            Load Data
                          </CButton>
                          &nbsp;&nbsp; Show All :
                          <CSwitch
                            color="success"
                            v-bind="{ shape: 'pill' }"
                            :checked.sync="lab_showall"
                          />
                          &nbsp;&nbsp;<CButton
                            color="behance"
                            @click="
                              () => {
                                this.insertExternalLabModal = true;
                              }
                            "
                          >
                            Insert External Lab Results
                          </CButton>
                           Legend: 🟩 - Internal Results  🟦 - External Results
                        </CCol>
                        <CCol v-if="lab_showall === true" style="padding: 5px" class="d-flex" align="right">
                          <CInput
                              placeholder="Enter Testcode to Filter"
                              id="filter"
                              v-model="lab_filter"
                              @input=" () => {
                                this.GetLabInvestigationData(this);
                              }"
                          />
                        </CCol>
                        <CDataTable
                          style="
                            padding: 5px;
                            text-align: left;
                            margin-top: 8px;
                            height: 100%;
                          "
                          header
                          clickableRows
                          :tableProps="{
                            striped: true,
                            hover: true,
                            border: true,
                          }"
                          border
                          hover
                          striped
                          :items="LabInvestData"
                          :fields="LabInvestfields"
                          :items-per-page="40"
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
                        <CCardBody>
                          <CRow align="center" style="margin-top: -1px">
                            <CCol col="5" style="">
                              <!-- <CButton @click="ad">TestTop</CButton> -->
                              <CSelect
                                class="align-self-left"
                                :value.sync="Prescript.drugsSearch"
                                :options="drugs"
                                @change="DisabledForPrescriptionOtherField"
                                style="width: 100%"
                              >
                              </CSelect>

                              <!-- <CDropdown  toggler-text="Drugs" class="m-6" style="width:100%">
                                    <CDropdownItem>Lab Investigations</CDropdownItem>
                                    <CDropdownItem>Othe Investigations</CDropdownItem>
                              </CDropdown> -->
                            </CCol>

                            <CCol col="5" align="left" style="">
                              <div>
                                <CRow>
                                  <CLable>Urgent &nbsp;</CLable>
                                  <CInputCheckbox
                                    type="checkbox"
                                    id="Urgent"
                                    :checked.sync="Prescript.urgentCheck"
                                  />
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

                          <CRow align="center" style="margin-top: -15px">
                            <CCol col="5" style="margin-left: 0px">
                              <CInput
                                list="drugs"
                                placeholder="Input your search"
                                v-model="Prescript.InputSearch"
                                @input="onDrugSearch"
                                @change="onDrugSelect"
                              />

                              <datalist id="drugs" loc>
                                <option
                                  v-for="drug in InputSearchSelect"
                                  :key="drug.label"
                                  :value="drug.value"
                                >
                                  {{ drug.label }}
                                </option>
                              </datalist>
                            </CCol>
                            <CCol col="2" style="margin-left: -25px">
                              <CInput
                                placeholder="Strenth"
                                id="Strenth"
                                v-model="Prescript.strenth"
                              />
                            </CCol>
                            <CCol
                              col="3"
                              align="left"
                              style="margin-left: -24px"
                            >
                              <!-- {{Prescript.selUnit}} -->
                              <CSelect
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
                              <CSelect
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
                              <CInput
                                placeholder="Comment"
                                v-model="Prescript.comment"
                              />
                            </CCol>

                            <CCol
                              col="4"
                              style="margin-left: -24px; margin-top: -13px"
                              align="left"
                            >
                              <CSelect
                                class="align-self-left"
                                id="Duration"
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
                              <CButton
                                color="info"
                                style="width: 8rem"
                                @click="AddForPrescription"
                              >
                                {{
                                  this.Prescript.index === null
                                    ? "Add"
                                    : "Update"
                                }}</CButton
                              >
                            </CCol>
                          </CRow>
                          <!-- <CCardHeader><strong> Urgent Prescriptions: </strong> </CCardHeader> -->
                          <!-- This Table for prescription  -->
                          <CRow style="margin-left: 0px; margin-top: -11px">
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
                                        <i @click="editPrescription(item)"
                                          ><CIcon
                                            name="cil-pencil"
                                            size="sm"
                                            class="ml-2 mb-2 text-primary"
                                        /></i>

                                        <i
                                          @click="
                                            deletePrescriptionModal = true;
                                            deletePrescriptionItem = item;
                                          "
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
                                        <i @click="editPrescription(item)"
                                          ><CIcon
                                            name="cil-pencil"
                                            size="sm"
                                            class="ml-2 mb-2 text-primary"
                                        /></i>

                                        <i
                                          @click="
                                            deleteDiagnosisModal = true;
                                            deletePrescriptionItem = item;
                                          "
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
                                style="margin-left: 0px; margin-right: -15px"
                              >
                                <CRow>
                                  <CCol col="3" class="d-flex">
                                    <CLable
                                      style="margin-top: 5px; margin-left: 20px"
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
                                      style="margin-top: 5px; margin-left: 10px"
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
                        <CCardBody style="margin-left: -9px; margin-top: -20px">
                          <!-- </CCol> -->
                        </CCardBody>
                        <CRow style="margin-top: -20px; padding: 5px">
                          <CCol col="9">
                            <form>
                              <!-- * Here the firaz said to unable to validation  tue-08-02-2022-->
                              <CTextarea
                                label="Comment (Print)"
                                valid-feedback="Input is valid."
                                invalid-feedback="Please provide comment(Print)."
                                value="Valid value"
                                v-model="header.comment_print"
                                rows="8"
                                cols="50"
                              />
                              <CTextarea
                                label="Plan"
                                v-model="header.comment_internal"
                                rows="8"
                                cols="50"
                              />
                            </form>
                          </CCol>
                          <!-- &nbsp;&nbsp; -->
                          <CCol col="3" style="margin-left: -25px">
                            <!-- <h6 @click="modaltest()">Click My Tesk</h6> -->
                            <CCard style="margin-right: -30px; margin-top: 5px">
                              <div class="custom-file">
                                <div @click="uploadImage">
                                  <input
                                    type=""
                                    class="custom-file-input"
                                    style="cursor: pointer"
                                    id="customFile"
                                    @click="uploadImage"
                                  />
                                  <label
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
                                  @click="imageload(data)"
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
                  <CTab title="Referral Form">
                    <CTabContent style="margin-top: 50px; padding: 0px">
                      <CTabPane v-if="tabIndex === 4">
                        <CRow>
                          <CCol class="" style="margin-left: 2px">
                            <CInput
                              list="Dr"
                              placeholder="Input your search"
                              v-model="ReferralForm.Doctor"
                              @input="SearchDoctor()"
                              @change="GetCodeOfDoct()"
                            />

                            <datalist id="Dr">
                              <option
                                v-for="(data,idx) in DoctorNameFilter"
                                :key="data.value + idx"
                                :value="data.label"
                              >
                                {{ data.label }}
                              </option>
                            </datalist>
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
                            <CButton
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
                            <CButton
                              color="tumblr"
                              style="
                                width: 10rem;
                                margin-right: 30px;
                                height: 2.3rem;
                              "
                              @click="UpdateasReferralform"
                            >
                              Update as Referral
                            </CButton>
                            <CButton
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
                              align="right"
                              @click="loadTemplate"
                            >
                              Load Template
                            </CButton>
                            <!-- {{ReferralLetterform}} -->
                            <CSelect
                              class="col align-self-right"
                              :value.sync="LetterID"
                              :options="ReferralLetter"
                              @change="GetReferralbyLetterId()"
                              style="
                                margin-left: -10px;
                                margin-right: -15px;
                                width: 20%;
                              "
                            >
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
                        <form style="margin-left: 10px">
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
                        </form>
                      </CTabPane>
                    </CTabContent>
                  </CTab>
                  <CTab title="Checked list">
                    <CTabContent style="margin-top: 50px; padding: 0px">
                      <CTabPane>
                        <CRow
                          md="9"
                          v-for="item in checkedList"
                          style="margin-left: 5px"
                          :key="item.ItemCode"
                          align="left"
                        >
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
                        <div class="align_set" style="margin-left: 10px">
                          <CButton
                            color="behance"
                            style="width: 8rem"
                            @click="onCheckListSave"
                          >
                            Save
                          </CButton>
                          &nbsp;
                          <CButton color="twitter" style="width: 8rem">
                            Clear
                          </CButton>
                        </div>
                        <CDataTable
                          class=""
                          style="
                            margin-left: 10px;
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
                  <CTab title="Pregnancy Diabetes" v-show="clinic !== 'MC' && false">
                    <CTabContent style="margin-top: 50px; padding: 0px">
                      <CTabPane>
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
                  <CTab title="Footcare" v-show="clinic !== 'MC' && false">
                    <CTabContent style="margin-top: 50px; padding: 0px">
                      <CTabPane>
                         <div class="row">
                                    <!--  ** Main Tab footcare Page 1 Start **  -->
                                    <div class="col-12">
                                        <div class="row">
                                            <div class="col-4">
                                                <div class="row">
                                                    <div class="col-6">
                                                        <div class="row">
                                                            <div class="col-6">
                                                                <label>DM Type</label>
                                                            </div>
                                                            <div class="col-6">
                                                                <div class="input-group input-group-sm">
                                                                    <input type="text" id="dmType" class="form-control" placeholder="DM Type">
                                                                </div>
                                                            </div>
                                                        </div>
                                                    </div>
                                                    <div class="col-1"> </div>
                                                    <div class="col-5">
                                                        <div class="row">
                                                            <div class="col-6">
                                                                <label>Duration</label>
                                                            </div>
                                                            <div class="col-6">
                                                                <div class="input-group input-group-sm">
                                                                    <input type="text" id="duration" class="form-control" placeholder="In Year">
                                                                </div>
                                                            </div>
                                                        </div>

                                                    </div>
                                                </div>
                                            </div>
                                            <div class="col-1"> </div>
                                            <div class="col-3">
                                                <div class="row">
                                                    <div class="col-4">
                                                        <label>Treatment</label>
                                                    </div>
                                                    <div class="col-1"> </div>

                                                    <div class="col-7">
                                                        <div class="input-group input-group-sm">
                                                            <select class="form-control select2bs4" id="fcTreatment" data-placeholder="Treatment" style="width: 100%;">
                                                                <option selected="selected">None</option>
                                                                <option>OHD</option>
                                                                <option>Insulin</option>
                                                            </select>
                                                        </div>

                                                    </div>
                                                </div>
                                            </div>
                                            <div class="col-1"> </div>

                                            <div class="col-3">
                                                <div class="row">
                                                    <div class="col-4">
                                                        <label>HbA1c</label>
                                                    </div>
                                                    <div class="col-1"> </div>

                                                    <div class="col-7">
                                                        <div class="input-group input-group-sm">
                                                            <input type="text" id="hbA1c" class="form-control" placeholder="HbA1c">
                                                        </div>

                                                    </div>
                                                </div>

                                            </div>
                                        </div>   <!-- First Row -->

                                        <div class="row">
                                            <div class="col-4">
                                                <div class="row">
                                                    <div class="col-5">
                                                        <label>Impaired Vision</label>
                                                    </div>
                                                    <div class="col-1"> </div>

                                                    <div class="col-6">
                                                        <div class="input-group input-group-sm">
                                                            <select id="ImpairedVersion" class="form-control select2bs4" data-placeholder="Impaired Vision" style="width: 100%;">
                                                                <option>IHD</option>
                                                                <option>HT</option>
                                                                <option>CKD</option>
                                                                <option>CVA</option>
                                                                <option>PVD</option>
                                                            </select>
                                                        </div>

                                                    </div>
                                                </div>
                                            </div>
                                            <div class="col-1"> </div>
                                            <div class="col-3">
                                                <div class="row">
                                                    <div class="col-4">
                                                        <label>Smoking</label>
                                                    </div>
                                                    <div class="col-1"> </div>

                                                    <div class="col-7">
                                                        <div class="input-group input-group-sm">
                                                            <select id="fcSmoking" class="form-control select2bs4" data-placeholder="Smoking" style="width: 100%;">
                                                                <option selected="selected">None</option>
                                                                <option>Present</option>
                                                                <option>Past</option>
                                                            </select>
                                                        </div>

                                                    </div>
                                                </div>
                                            </div>
                                            <div class="col-1"> </div>

                                            <div class="col-3">
                                                <div class="row">
                                                    <div class="col-4">
                                                        <label>Other</label>
                                                    </div>
                                                    <div class="col-1"> </div>

                                                    <div class="col-7">
                                                        <div class="input-group input-group-sm">
                                                            <input type="text" id="fcOther" class="form-control" placeholder="Other">
                                                        </div>

                                                    </div>
                                                </div>

                                            </div>
                                        </div>   <!-- Second Row -->

                                        <div class="row">
                                            <div class="col-12 callout callout-success">
                                                <p>General Assessment</p>
                                            </div>
                                        </div>   <!--General Assessment-->
                                        <div class="row"></div>
                                        <!--  ** R L Skin and Nails Section Start **  -->
                                        <div class="row">
                                            <div class="col-4" style="text-align:left">
                                                <h3>Skin and Nails</h3>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <strong>RIGHT</strong>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <strong>LEFT</strong>
                                            </div>
                                            <div class="col-4">
                                                <strong></strong>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <strong>RIGHT</strong>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <strong>LEFT</strong>
                                            </div>

                                        </div>      <!--Skin and Nails-->

                                        <hr>
                                        <div class="row">
                                            <div class="col-4" style="text-align:left">
                                                <label>Dry Skin</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="DrySkinRight">
                                                </div>

                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="DrySkinLeft">
                                                </div>
                                            </div>
                                            <div class="col-4">
                                                <label>Web Space Infection</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="WebSpaceInfectionRight">
                                                </div>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="WebSpaceInfectionLeft">
                                                </div>
                                            </div>
                                        </div>

                                        <div class="row">
                                            <div class="col-4" style="text-align:left">
                                                <label>Callus/ Corns</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="CallusCornsRight">
                                                </div>

                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="CallusCornsLeft">
                                                </div>
                                            </div>
                                            <div class="col-4">
                                                <label>Nail Bed Infection</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="NailBedInfectionRight">
                                                </div>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="NailBedInfectionLeft">
                                                </div>
                                            </div>
                                        </div>
                                        <div class="row">
                                            <div class="col-4" style="text-align:left">
                                                <label>Fissures / Cracks</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="FissuresCracksRight">
                                                </div>

                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="FissuresCracksLeft">
                                                </div>
                                            </div>
                                            <div class="col-4">
                                                <label>In Growing Toe Nails</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="InGrowingToeNailsRight">
                                                </div>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="InGrowingToeNailsLeft">
                                                </div>
                                            </div>
                                        </div>

                                        <div class="row">
                                            <div class="col-12 callout callout-success">
                                                <p>Risk Category</p>
                                            </div>
                                        </div>   <!--Risk Category-->
                                        <div class="row">
                                            <div class="col-6" style="text-align:left">
                                                <h3>Deformity (Any 1 of the following) </h3>
                                            </div>
                                        </div>
                                        <hr>

                                        <div class="row">
                                            <div class="col-4" style="text-align:left">
                                                <label>Hammer Toes</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="HammerToesRight">
                                                </div>

                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="HammerToesLeft">
                                                </div>
                                            </div>
                                            <div class="col-4">
                                                <label>Reduced Ankle Reflex</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="ReducedAnkleReflexRight">
                                                </div>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="ReducedAnkleReflexLeft">
                                                </div>
                                            </div>
                                        </div>

                                        <div class="row">
                                            <div class="col-6">
                                                <div class="row">
                                                    <div class="col-8" style="text-align:left">
                                                        <label>Claw toes</label>
                                                    </div>
                                                    <div class="col-2" style="text-align:center">
                                                        <div class="icheck-primary d-inline">
                                                            <input type="checkbox" id="ClawtoesRight">
                                                        </div>

                                                    </div>
                                                    <div class="col-2" style="text-align:center">
                                                        <div class="icheck-primary d-inline">
                                                            <input type="checkbox" id="ClawtoesLeft">
                                                        </div>
                                                    </div>
                                                </div>
                                                <div class="row">
                                                    <div class="col-8" style="text-align:left">
                                                        <label>Overlapping digits</label>
                                                    </div>
                                                    <div class="col-2" style="text-align:center">
                                                        <div class="icheck-primary d-inline">
                                                            <input type="checkbox" id="OLDRight">
                                                        </div>

                                                    </div>
                                                    <div class="col-2" style="text-align:center">
                                                        <div class="icheck-primary d-inline">
                                                            <input type="checkbox" id="OLDLeft">
                                                        </div>
                                                    </div>
                                                </div>

                                            </div>
                                            <div class="col-6">
                                                <div class="row">
                                                    <div class="col-8">
                                                        <label>Positive Monofilament Test-if Unable to feel less than 8-(+)</label>
                                                    </div>
                                                    <div class="col-2" style="text-align:center">
                                                        <div class="icheck-primary d-inline">
                                                            <input type="checkbox" id="PositiveMonofilamentRight">
                                                        </div>
                                                    </div>
                                                    <div class="col-2" style="text-align:center">
                                                        <div class="icheck-primary d-inline">
                                                            <input type="checkbox" id="PositiveMonofilamentLeft">
                                                        </div>
                                                    </div>
                                                </div>
                                            </div>
                                        </div>

                                        <div class="row">
                                            <div class="col-6">
                                                <div class="row">
                                                    <div class="col-8" style="text-align:left">
                                                        <label>Bunion</label>
                                                    </div>
                                                    <div class="col-2" style="text-align:center">
                                                        <div class="icheck-primary d-inline">
                                                            <input type="checkbox" id="BunionRight">
                                                        </div>

                                                    </div>
                                                    <div class="col-2" style="text-align:center">
                                                        <div class="icheck-primary d-inline">
                                                            <input type="checkbox" id="BunionLeft">
                                                        </div>
                                                    </div>
                                                </div>
                                                <div class="row">
                                                    <div class="col-8" style="text-align:left">
                                                        <label>Arch deformities</label>
                                                    </div>
                                                    <div class="col-2" style="text-align:center">
                                                        <div class="icheck-primary d-inline">
                                                            <input type="checkbox" id="ADRight">
                                                        </div>

                                                    </div>
                                                    <div class="col-2" style="text-align:center">
                                                        <div class="icheck-primary d-inline">
                                                            <input type="checkbox" id="ADLeft">
                                                        </div>
                                                    </div>
                                                </div>

                                            </div>
                                            <div class="col-6">
                                                <div class="row">
                                                    <div class="col-8">
                                                        <label>Abnormal Biothesiometer Test Loss of protective sensation</label>
                                                    </div>
                                                    <div class="col-2" style="text-align:center">
                                                        <div class="icheck-primary d-inline">
                                                            <input type="checkbox" id="ABTRight">
                                                        </div>
                                                    </div>
                                                    <div class="col-2" style="text-align:center">
                                                        <div class="icheck-primary d-inline">
                                                            <input type="checkbox" id="ABTLeft">
                                                        </div>
                                                    </div>
                                                </div>
                                            </div>
                                        </div>

                                        <div class="row">
                                            <div class="col-6" style="text-align:left">
                                                <h3>Vasoulopathy (Palpable D & PT)</h3>
                                            </div>

                                            <div class="col-4">
                                                <h3>Other</h3>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <strong></strong>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <strong></strong>
                                            </div>

                                        </div>      <!--Vasoulopathy (Palpable D & PT)-->
                                        <hr>
                                        <div class="row">
                                            <div class="col-4" style="text-align:left">
                                                <label>Absent Dorsalis pedis</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="ADPRight">
                                                </div>

                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="ADPLeft">
                                                </div>
                                            </div>
                                            <div class="col-4">
                                                <label>Previous ulceration</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="PreviousUlcerationRight">
                                                </div>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="PreviousUlcerationLeft">
                                                </div>
                                            </div>
                                        </div>

                                        <div class="row">
                                            <div class="col-4" style="text-align:left">
                                                <label>Absent Posterior tibial</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="APTRight">
                                                </div>

                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="APTLeft">
                                                </div>
                                            </div>
                                            <div class="col-4">
                                                <label>Previous amputation</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="PreviousAmputationRight">
                                                </div>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="PreviousAmputationLeft">
                                                </div>
                                            </div>
                                        </div>

                                        <div class="row">
                                            <div class="col-4" style="text-align:left">
                                                <label>ABPL</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="input-group input-group-sm">
                                                    <input type="text" id="abplR" class="form-control" placeholder="ABPL">
                                                </div>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="input-group input-group-sm">
                                                    <input type="text" id="abplL" class="form-control" placeholder="ABPL">
                                                </div>
                                            </div>
                                            <div class="col-2">
                                                <label>Specify</label>
                                            </div>
                                            <div class="col-4" style="text-align:center">
                                                <div class="input-group input-group-sm">
                                                    <input type="text" id="Specify" class="form-control" placeholder="Specify">
                                                </div>
                                            </div>
                                        </div>


                                        <div class="row">
                                            <div class="col-4" style="text-align:left">
                                                <label>CLI (IF ABPL =< 0.5)</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="IFAPTRight">
                                                </div>

                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="IFAPTLeft">
                                                </div>
                                            </div>
                                            <div class="col-4">
                                                <label>On renal replacement therapy (Y/N)</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="ORRTRight">
                                                </div>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="ORRTLeft">
                                                </div>
                                            </div>
                                        </div>

                                        <hr>
                                        <div class="row">
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="VasAnuR">
                                                    <label>Right</label>
                                                </div>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="VasAnuL">
                                                    <label>Left</label>
                                                </div>
                                            </div>
                                            <div class="col-2">
                                                <label>Low Risk</label>
                                            </div>
                                            <div class="col-5">
                                                <label>No Risk factors present expect callus alone</label>
                                            </div>
                                            <div class="col-3">
                                                <label>Annual following up</label>
                                            </div>
                                        </div>
                                        <hr>
                                        <div class="row">
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="VasSixR">
                                                    <label>Right</label>
                                                </div>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="VasSixL">
                                                    <label>Left</label>
                                                </div>
                                            </div>
                                            <div class="col-2">
                                                <label>Moderate Risk</label>
                                            </div>
                                            <div class="col-5">
                                                <div class="row">
                                                    <label>Deformity or</label>
                                                </div>
                                                <div class="row">
                                                    <label>Neuropathy or</label>
                                                </div>
                                                <div class="row">
                                                    <label>Non critical limb ischemia</label>
                                                </div>

                                            </div>
                                            <div class="col-3">
                                                <label>6 Months following up</label>
                                            </div>
                                        </div>
                                        <hr>
                                        <div class="row">
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="VasThreeR">
                                                    <label>Right</label>
                                                </div>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="VasThreeL">
                                                    <label>Left</label>
                                                </div>
                                            </div>
                                            <div class="col-2">
                                                <label>High Risk</label>
                                            </div>
                                            <div class="col-5">
                                                <div class="row">
                                                    <label>Previous ulceration or</label>
                                                </div>
                                                <div class="row">
                                                    <label>Previous amputation or</label>
                                                </div>
                                                <div class="row">
                                                    <label>On renal replacement therapy or</label>
                                                </div>
                                                <div class="row">
                                                    <label>Non critical limb ischemia in combination with callus and / or deformity</label>
                                                </div>

                                            </div>
                                            <div class="col-3">
                                                <label>3 Months following up</label>
                                            </div>
                                        </div>
                                        <hr>

                                        <div class="row">
                                            <div class="col-12 callout callout-success">
                                                <p>Emergency acute foot conditions</p>
                                            </div>
                                        </div>   <!--Emergency acute foot conditions-->

                                        <div class="row">
                                            <div class="col-6" style="text-align:left">
                                                <h3>Acute Diabetic Foot</h3>
                                            </div>

                                            <div class="col-4">
                                                <h3></h3>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <strong></strong>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <strong></strong>
                                            </div>

                                        </div>      <!--Acute Diabetic Foot-->

                                        <div class="row">
                                            <div class="col-4" style="text-align:left">
                                                <label>Cellulites</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="CellulitesR">
                                                </div>

                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="CellulitesL">
                                                </div>
                                            </div>
                                            <div class="col-4">
                                                <label>Gangrene</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="GangreneR">
                                                </div>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="GangreneL">
                                                </div>
                                            </div>
                                        </div>

                                        <div class="row">
                                            <div class="col-4" style="text-align:left">
                                                <label>Acute Ulcer</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="AcuteUlcerR">
                                                </div>

                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="AcuteUlcerL">
                                                </div>
                                            </div>
                                            <div class="col-4">
                                                <label>Acute Charcot</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="AcuteCharcotR">
                                                </div>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="AcuteCharcotL">
                                                </div>
                                            </div>
                                        </div>

                                        <div class="row">
                                            <div class="col-4" style="text-align:left">
                                                <label>Sepsis</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="SepsisR">
                                                </div>

                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="SepsisL">
                                                </div>
                                            </div>
                                            <div class="col-4">
                                                <label>Other</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="ADFOtherR">
                                                </div>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="ADFOtherL">
                                                </div>
                                            </div>
                                        </div>



                                        <div class="row">
                                            <div class="col-6" style="text-align:left">
                                                <h3>Foot care and Footwear</h3>
                                            </div>

                                            <div class="col-4">
                                                <h3></h3>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <strong></strong>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <strong></strong>
                                            </div>

                                        </div>      <!--Foot care and Footwear-->

                                        <div class="row">
                                            <div class="col-4" style="text-align:left">
                                                <label>Satisfactory foot hygiene</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="SFHRight">
                                                </div>

                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="SFHLeft">
                                                </div>
                                            </div>
                                            <div class="col-4">
                                                <label>Appropriate</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="AppropriateR">
                                                </div>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="AppropriateL">
                                                </div>
                                            </div>
                                        </div>

                                        <div class="row">
                                            <div class="col-4" style="text-align:left">
                                                <label>Education Received</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="EducationReceivedR">
                                                </div>

                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="EducationReceivedL">
                                                </div>
                                            </div>
                                            <div class="col-4">
                                                <label>Normal</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="NormalR">
                                                </div>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="NormalL">
                                                </div>
                                            </div>
                                        </div>

                                        <div class="row">
                                            <div class="col-4" style="text-align:left">
                                                <label>Satisfactory adherence</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="SARight">
                                                </div>

                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="SALeft">
                                                </div>
                                            </div>
                                            <div class="col-4">
                                                <label>Diabetic Shoe</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="DiabeticShoeR">
                                                </div>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="DiabeticShoeL">
                                                </div>
                                            </div>
                                        </div>

                                        <div class="row">
                                            <div class="col-4" style="text-align:left">
                                                <label></label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                            </div>
                                            <div class="col-4">
                                                <label>Thereapeutic Shoe</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="ThereapeuticShoeR">
                                                </div>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="ThereapeuticShoeL">
                                                </div>
                                            </div>
                                        </div>


                                        <div class="row">
                                            <div class="col-6" style="text-align:left">
                                                <h3>Referrals & Treatment</h3>
                                            </div>

                                            <div class="col-4">
                                                <h3></h3>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <strong></strong>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <strong></strong>
                                            </div>

                                        </div>      <!--Referrals & Treatmentr-->

                                        <div class="row">
                                            <div class="col-4" style="text-align:left">
                                                <label>Debridement of callus</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="DebridementOfCallus">
                                                </div>
                                            </div>
                                            <div class="col-1"></div>
                                            <div class="col-4" style="text-align:left">
                                                <label>Appropriate Footwear</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="AppropriateFootwear">
                                                </div>
                                            </div>
                                            <div class="col-1"></div>
                                        </div>

                                        <div class="row">
                                            <div class="col-4" style="text-align:left">
                                                <label>Offloading Shoe</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="OffloadingShoe">
                                                </div>
                                            </div>
                                            <div class="col-1"></div>
                                            <div class="col-4" style="text-align:left">
                                                <label>Vascular Clinic</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="VascularClinic">
                                                </div>
                                            </div>
                                            <div class="col-1"></div>
                                        </div>

                                        <div class="row">
                                            <div class="col-4" style="text-align:left">
                                                <label>Medication</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="Medication">
                                                </div>
                                            </div>
                                            <div class="col-1"></div>
                                            <div class="col-4" style="text-align:left">
                                                <label>Ulcer Clinic</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="UlcerClinic">
                                                </div>
                                            </div>
                                            <div class="col-1"></div>
                                        </div>

                                        <div class="row">
                                            <div class="col-4" style="text-align:left">
                                                <label>Education</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="Education">
                                                </div>
                                            </div>
                                            <div class="col-1"></div>
                                            <div class="col-4" style="text-align:left">
                                                <label>Orthotist</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="Orthotist">
                                                </div>
                                            </div>
                                            <div class="col-1"></div>
                                        </div>

                                        <div class="row">
                                            <div class="col-4" style="text-align:left">
                                                <label>Physiotherapy</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="Physiotherapy">
                                                </div>
                                            </div>
                                            <div class="col-1"></div>
                                            <div class="col-4" style="text-align:left">
                                                <label>Other</label>
                                            </div>
                                            <div class="col-1" style="text-align:center">
                                                <div class="icheck-primary d-inline">
                                                    <input type="checkbox" id="RFOther">
                                                </div>
                                            </div>
                                            <div class="col-1"></div>
                                        </div>
                                        <div class="row">
                                            <div class="col-12 callout callout-success">
                                                <p>Foot Scan</p>
                                            </div>
                                        </div>
                                        <div class="row">
                                            <div class="col-md-8">
                                                <div class="bg-img-right-feet-pointer">
                                                    <form class="container">
                                                        <div class="btn-right-feet-one">
                                                            <div class="icheck-primary d-inline">
                                                                <input type="checkbox" id="rightPointerOne">
                                                            </div>
                                                        </div>
                                                        <div class="btn-right-feet-two">
                                                            <div class="icheck-primary d-inline">
                                                                <input type="checkbox" id="rightPointerTwo">
                                                            </div>
                                                        </div>
                                                        <div class="btn-right-feet-three">
                                                            <div class="icheck-primary d-inline">
                                                                <input type="checkbox" id="rightPointerThree">
                                                            </div>
                                                        </div>
                                                        <div class="btn-right-feet-four">
                                                            <div class="icheck-primary d-inline">
                                                                <input type="checkbox" id="rightPointerFour">
                                                            </div>
                                                        </div>
                                                        <div class="btn-right-feet-five">
                                                            <div class="icheck-primary d-inline">
                                                                <input type="checkbox" id="rightPointerFive">
                                                            </div>
                                                        </div>
                                                        <div class="btn-right-feet-six">
                                                            <div class="icheck-primary d-inline">
                                                                <input type="checkbox" id="rightPointerSix">
                                                            </div>
                                                        </div>
                                                        <div class="btn-right-feet-seven">
                                                            <div class="icheck-primary d-inline">
                                                                <input type="checkbox" id="rightPointerSeven">
                                                            </div>
                                                        </div>
                                                        <div class="btn-right-feet-eight">
                                                            <div class="icheck-primary d-inline">
                                                                <input type="checkbox" id="rightPointerEight">
                                                            </div>
                                                        </div>
                                                        <div class="btn-right-feet-nine">
                                                            <div class="icheck-primary d-inline">
                                                                <input type="checkbox" id="rightPointerNine">
                                                            </div>
                                                        </div>
                                                        <div class="btn-right-feet-ten">
                                                            <div class="icheck-primary d-inline">
                                                                <input type="checkbox" id="rightPointerTen">
                                                            </div>
                                                        </div>
                                                    </form>
                                                </div>
                                            </div>
                                            <div class="col-md-4">
                                                <div class="bg-img-right-feet-values">
                                                    <form class="container">
                                                        <div class="textRightPointOne">
                                                            <input type="text" id="rightPointOneValue" style="width: 30px;" />
                                                        </div>
                                                        <div class="textRightPointTwo">
                                                            <input type="text" id="rightPointTwoValue" style="width: 30px;" />
                                                        </div>
                                                        <div class="textRightPointThree">
                                                            <input type="text" id="rightPointThreeValue" style="width: 30px;" />
                                                        </div>
                                                        <div class="textRightPointFour">
                                                            <input type="text" id="rightPointFourValue" style="width: 30px;" />
                                                        </div>
                                                        <div class="textRightPointFive">
                                                            <input type="text" id="rightPointFiveValue" style="width: 30px;" />
                                                        </div>
                                                        <div class="textRightPointSix">
                                                            <input type="text" id="rightPointSixValue" style="width: 30px;" />
                                                        </div>
                                                    </form>
                                                </div>
                                            </div>
                                        </div>
                                        <br />
                                        <div class="row">
                                            <div class="col-md-8">
                                                <div class="bg-img-left-feet-pointer">
                                                    <form class="container">
                                                        <div class="btn-left-feet-one">
                                                            <div class="icheck-primary d-inline">
                                                                <input type="checkbox" id="leftPointerOne">
                                                            </div>
                                                        </div>
                                                        <div class="btn-left-feet-two">
                                                            <div class="icheck-primary d-inline">
                                                                <input type="checkbox" id="leftPointerTwo">
                                                            </div>
                                                        </div>
                                                        <div class="btn-left-feet-three">
                                                            <div class="icheck-primary d-inline">
                                                                <input type="checkbox" id="leftPointerThree">
                                                            </div>
                                                        </div>
                                                        <div class="btn-left-feet-four">
                                                            <div class="icheck-primary d-inline">
                                                                <input type="checkbox" id="leftPointerFour">
                                                            </div>
                                                        </div>
                                                        <div class="btn-left-feet-five">
                                                            <div class="icheck-primary d-inline">
                                                                <input type="checkbox" id="leftPointerFive">
                                                            </div>
                                                        </div>
                                                        <div class="btn-left-feet-six">
                                                            <div class="icheck-primary d-inline">
                                                                <input type="checkbox" id="leftPointerSix">
                                                            </div>
                                                        </div>
                                                        <div class="btn-left-feet-seven">
                                                            <div class="icheck-primary d-inline">
                                                                <input type="checkbox" id="leftPointerSeven">
                                                            </div>
                                                        </div>
                                                        <div class="btn-left-feet-eight">
                                                            <div class="icheck-primary d-inline">
                                                                <input type="checkbox" id="leftPointerEight">
                                                            </div>
                                                        </div>
                                                        <div class="btn-left-feet-nine">
                                                            <div class="icheck-primary d-inline">
                                                                <input type="checkbox" id="leftPointerNine">
                                                            </div>
                                                        </div>
                                                        <div class="btn-left-feet-ten">
                                                            <div class="icheck-primary d-inline">
                                                                <input type="checkbox" id="leftPointerTen">
                                                            </div>
                                                        </div>
                                                    </form>
                                                </div>
                                            </div>
                                            <div class="col-md-4">
                                                <div class="bg-img-left-feet-values">
                                                    <form class="container">
                                                        <div class="textLeftPointOne">
                                                            <input type="text" id="leftPointOneValue" style="width: 30px;" />
                                                        </div>
                                                        <div class="textLeftPointTwo">
                                                            <input type="text" id="leftPointTwoValue" style="width: 30px;" />
                                                        </div>
                                                        <div class="textLeftPointThree">
                                                            <input type="text" id="leftPointThreeValue" style="width: 30px;" />
                                                        </div>
                                                        <div class="textLeftPointFour">
                                                            <input type="text" id="leftPointFourValue" style="width: 30px;" />
                                                        </div>
                                                        <div class="textLeftPointFive">
                                                            <input type="text" id="leftPointFiveValue" style="width: 30px;" />
                                                        </div>
                                                        <div class="textLeftPointSix">
                                                            <input type="text" id="leftPointSixValue" style="width: 30px;" />
                                                        </div>
                                                    </form>
                                                </div>
                                            </div>
                                        </div>


                                        <!--  ** R L Skin and Nails Section End **  -->
                                    </div>
                                    <!--  ** Main Tab footcare Page 1 End **  -->
                                </div>
                      </CTabPane>
                    </CTabContent>
                  </CTab>
                </CTabs>
              </CCol>

              <CCol>
                <CCardBody style="margin-left: -78px; height: 100%">
                  <CCol>
                    <CButton color="success" @click="onSaveClick"
                      >Save
                    </CButton>
                  </CCol>
                  <CCol>
                    <CDropdown toggler-text="Print" color="yahoo">
                      <CDropdownItem disabled>Footcare</CDropdownItem>
                      <CDropdownItem>Initial visit</CDropdownItem>
                      <CDropdownItem @click="onShowReferralPDF" >Referral form</CDropdownItem>
                      <CDropdownItem @click="onShowSummaryPDF"
                        >Summary</CDropdownItem
                      >
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
import { DateTime } from "luxon";
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
      vf: null,
      // here for Referral
      ReferraltemplateName: false,
      ReferralTemp: null,
      ReferralLetterform: null,
      //Final Save
      ReferralLetterTemplateArray: null,
      PatientReferralbyID: {},
      //
      alertboxSuccess: false,
      alertboxFailed: false,
      deletePrescriptionModal: false,
      insertExternalLabModal: false,
      insertNewProblemModal: false,
      deletePrescriptionItem: null,
      deleteDiagnosisModal: false,
      deleteDiagnosisItem: null,
      ClinicSession: null,
      NewSession: [{ label: "New", value: "--" }],

      ClinicWiseSession: null,

      DocumentHeaderNotes: null,
      problemList: null,
      GetDoctorNamelist: null,
      DoctorValue: null,
      GetDocumentHeaderImagelist: null,
      GetDosageList: null,
      GetUnitList: null,
      Profile: null,
      BodyMI: [],
      labAndDrug: null,
      DP: localStorage.getItem("DP"),
      ImageOfDisease: null,
      ProfileImage: null,
      imagePreviewFordisease: null,
      NewProblem: {
        description: null,
      },
      disease: {
        name: null,
        discriptions: null,
      },
      diseaseFinalSave: null,
      //
      DoctorNameFilter: null,
      Photo: null,
      // PRESCRIPTION MAIN ARRAY
      DataforPrescriptons: [],
      ImageOfDiseaseView: [],
      // FOR PRESCRIPTION TABLE
      DataforPrescriptonslist: [],

      // URGENT PRESCRIPTION TABLE
      UrgentPrescriptionDatalist: [],
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
      InputSearchSelect: null,
      tabIndex: 0,
      loadtemplate: false,
      mytext: false,
      warningModal: false,
      imagePreview: "/img/avatars/blank.png",
      AltImage: "Image.jpg",
      imageModal: false,
      alertModal: false,
      AlertMessage: "",
      lab_filter: null,
      //
      patient: null,
      InputSearch: null,
      Privilages: null,
      ShowProbAddOpt: false,
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
      ClinicType: [
        { label: "Medical Clinic", value: "MC" },
        { label: "Endocrine", value: "D&E" },
      ],
      //diagnosis
      diagnosis: {
        index: null,
        ProblemDescriptionLabel: null,
        ProblemDescriptionValue: null,
        ProblemNote: null,
        ProblemPriority: 1,
        ProblemStatus: true,
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
      clinic : String(localStorage.getItem("Location")),
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
      //
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
        { value: "1 Day", label: "1 Day" },
        { value: "2 Days", abel: "2 Days" },
        { value: "3 Days", label: "3 Days" },
        { value: "4 Days", label: "4 Days" },
        { value: "5 Days", label: "5 Days" },
        { value: "6 Days", label: "6 Days" },
        { value: "7 Days", label: "7 Days" },
        { value: "8 Days", label: "8 Days" },
        { value: "9 Days", label: "9 Days" },
        { value: "10 Days", label: "10 Days" },
        { value: "Daily", label: "Daily" },
        { value: "1 Week", label: "1 Week" },
        { value: "2 Weeks", label: "2 Weeks" },
        { value: "3 Weeks", label: "3 Weeks" },
        { value: "6 Weeks", label: "6 Weeks" },
        { value: "1 Month", label: "1 Month" },
        { value: "2 Months", label: "2 Months" },
        { value: "3 Months", label: "3 Months" },
        { value: "4 Months", label: "4 Months" },
        { value: "5 Months", label: "5 Months" },
        { value: "6 Months", label: "6 Months" },
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
      Template: [],
      checkedGridFields: [
        {
          filter: false,
          sorter: false,
          key: "DocumentID",
          label: "DocumentID",
        },
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
          _style: { width: "24%" },
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
          key: "show_details",
          label: "Action",
          _style: { width: "1%" },
          filter: false,
          sorter: false,
          _props: { color: "primary", className: "fw-semibold mr-20" },
        },
      ],
      checkedGrid: null,
      datey: ["2022-04-05", "2022-04-15"],
      checkedList: [],
      lab_showall: false,
      LabInvestfields: [
        {
          filter: false,
          sorter: false,
          key: "element_range",
          label: "Element Name / Range",
          _style: { width: "20%" },
        },
      ],

      LabInvestData: [],
      LabInvestInputData: [],
    };
  },

  //Validation
  // computed: {},
  //
  methods: {
    async onShowSummaryPDF() {
      let clinic = String(localStorage.getItem("Location"));
      const result = await this.api.get("/report/SummaryReport", {
        params: { DocID: String(localStorage.getItem("DDOCUMENT")) },
      });
      const result2 = await this.api.get("/username");
      let header = result.data.header;
      let consultant = result.data.consultant;
      let documentID = result.data.documentID;
      let prescriptionsSrc =
        result.data.prescriptions === undefined
          ? []
          : result.data.prescriptions;
      let problemListSrc =
        result.data.problemList === undefined ? [] : result.data.problemList;
      let profileData = this.Profile[0];
      // A4 dimensions : 210 x 297
      const doc = new this.$PDF();
      doc.setProperties({
        title: "Summary Report",
        subject: "Report",
        author: "admin",
        keywords: "",
        creator: "SJGH",
      });

      let img = new Image();
      img.src = "/sjgh/logo.png";
      doc.addImage(img, "png", 15, 10, 20, 20);
      let problemsList = [];
      let drugList = [];
      let investigationList = [];
      if (problemListSrc !== null) {
        problemListSrc.forEach((item, index, arr) => {
          problemsList.push([
            {
              content: `•	${
                item.ProblemDescription === null || item.ProblemDescription === "Others" ||
                item.ProblemDescription === undefined
                  ? ""
                  : item.ProblemDescription + "-"
              } ${
                item.Note === null || item.Note === undefined ? "" : item.Note
              }`,
              colSpan: 6,
            },
          ]);
        });
      }
      if (prescriptionsSrc !== null) {
        prescriptionsSrc.forEach((item, index, arr) => {
          if (!item.Urgent) {
            let ID = item.ID;
            let code = item.ItemCode;
            let comment =
              item.comment === null || item.comment === undefined
                ? ""
                : item.comment;
            let description = item.Description;
            let dosage = item.Dosage;
            let duration = item.duration;
            let strenth = item.Strenth;
            if (item.ItemType.trim() === "DRUG") {
              drugList.push([
                {
                  styles: { lineColor: [0, 0, 0] },
                  content: `${description === null ? "" : description}\n${
                    comment === null ? "" : comment
                  }`,
                  colSpan: 3,
                },
                {
                  styles: { lineColor: [0, 0, 0] },
                  content: `${strenth === null ? "" : strenth}`,
                },
                {
                  styles: { lineColor: [0, 0, 0] },
                  content: `${dosage === null ? "" : dosage}`,
                },
                {
                  styles: { lineColor: [0, 0, 0] },
                  content: `${duration === null ? "" : duration}`,
                },
              ]);
            } else {
              investigationList.push([
                {
                  content: `${
                    description === null
                      ? ""
                      : description.trim().toLowerCase().includes("other")
                      ? ""
                      : description + "-"
                  } ${comment === null ? "" : comment}`,
                  colSpan: 6,
                },
              ]);
            }
          }
        });
      }
      drugList.push([
        {
          content: "",
          colSpan: 3,
          styles: { lineWidth: 0 },
        },
        { styles: { lineWidth: 0 }, content: "" },
        { styles: { lineWidth: 0 }, content: "" },
        { styles: { lineWidth: 0 }, content: "" },
      ]);
      investigationList.push([
        {
          content: "",
          colSpan: 6,
        },
      ]);
      problemsList.push([
        {
          content: "",
          colSpan: 6,
        },
      ]);
      doc.autoTable({
        margin: { right: 0, left: 0 },
        styles: {
          fontStyle: "bold",
          fontSize: 16,
          cellPadding: 0.01,
        },
        columnStyles: {
          0: { halign: "center" },
        },
        theme: "plain",
        body: [
          [
            {
              content:
                clinic === "MC"
                  ? "Medical Clinic"
                  : "Endocrinology & Diabetes Centre",
              styles: { fontStyle: "bold", halign: "center" },
            },
          ],
          [{content:"Sri Jayewardenepura General Hospital",styles: { fontSize: 14, fontStyle: "normal", halign: "center" }}],
          [{content:"Thalapathpitiya, Nugegoda, Sri Lanka.",styles: { fontSize: 14,fontStyle: "normal",  halign: "center" }}],
          [
            {content:"Tel: 0112 778 610 Fax: 0112 778 611 Email: endocrine@sjghsrilanka.lk Web: sjgh.health.gov.lk	",styles: { fontSize: 12,fontStyle: "normal",  halign: "center" }},
          ],
          [""],
          ["Clinical Summary"]],
      });
      let planArray = []
      if(!(header.CommentInternal === null || header.CommentInternal.trim() === "")){
        planArray.push([{ content: "Plan", colSpan: 6, styles: { fontStyle: "bold" } }])
        planArray.push([{ content: header.CommentInternal, colSpan: 6 }] )
      }
      doc.autoTable({
        margin: { right: 20, left: 14 },
        styles: {
          fontStyle: "normal",
          cellPadding: 1,
          lineWidth: 0.01,
          lineColor: [255, 255, 255],
        },
        columnStyles: {
          0: { halign: "left", cellWidth: 30 },
          1: { halign: "left", cellWidth: 30  },
          2: { halign: "left", cellWidth: 30 },
          3: { halign: "left", cellWidth: 30 },
          4: { halign: "left", cellWidth: 30 },
          5: { halign: "left", cellWidth: 30 },
        },
        theme: "plain",
        body: [
          [
            {
              content: "Date",
              styles: { fontStyle: "bold", lineColor: [0, 0, 0] },
            },
            {
              styles: { lineColor: [0, 0, 0] },
              content: DateTime.fromISO(header.EntryDate.trim()).toFormat(
                "dd-MM-yyyy hh:mm"
              ),
            },
            {
              content: "Document No",
              styles: { fontStyle: "bold", lineColor: [0, 0, 0] },
            },
            {
              content: documentID.trim(),
              colSpan: 3,
              styles: { lineColor: [0, 0, 0] },
            },
          ],
          [
            {
              content: "Registration No ",
              styles: { fontStyle: "bold", lineColor: [0, 0, 0] },
            },
            {
              styles: { lineColor: [0, 0, 0] },
              content: header.RegistrationNo.trim(),
            },
            {
              content: "Age",
              styles: { fontStyle: "bold", lineColor: [0, 0, 0] },
            },
            {
              styles: { lineColor: [0, 0, 0] },
              colSpan: 3,
              content: profileData.Age,
            },
          ],
          [
            {
              content: "Patient Name",
              styles: { fontStyle: "bold", lineColor: [0, 0, 0] },
            },
            {
              styles: { lineColor: [0, 0, 0] },
              content: profileData.PatientName,
            },
            {
              content: "BP",
              styles: { fontStyle: "bold", lineColor: [0, 0, 0] },
            },
            { styles: { lineColor: [0, 0, 0] }, content: header.BP },
            {
              content: "Pulse Rate",
              styles: { fontStyle: "bold", lineColor: [0, 0, 0] },
            },
            { content: header.PulseRate, styles: { lineColor: [0, 0, 0] } },
          ],
          [{ content: "", colSpan: 6, styles: { lineWidth: 0 } }],
          [
            {
              content: "Problem List",
              colSpan: 6,
              styles: { fontStyle: "bold" },
            },
          ],
          ...problemsList,
          [{ content: "", colSpan: 6 }],
          [{ content: "Notes", colSpan: 6, styles: { fontStyle: "bold" } }],
          [{ content: header.Comment, colSpan: 6 }],
          ...planArray,
          [{ content: "", colSpan: 6 }],
          [
            {
              content: "Prescription",
              colSpan: 6,
              styles: { fontStyle: "bold" },
            },
          ],
          [
            {
              content: "Rx",
              colSpan: 3,
              styles: { fontStyle: "bold", lineColor: [0, 0, 0] },
            },
            {
              content: "Dosage",
              styles: { fontStyle: "bold", lineColor: [0, 0, 0] },
            },
            {
              content: "Frequency",
              styles: { fontStyle: "bold", lineColor: [0, 0, 0] },
            },
            {
              content: "Duration",
              styles: { fontStyle: "bold", lineColor: [0, 0, 0] },
            },
          ],
          ...drugList,
          [{ content: "", colSpan: 6, styles: { fontStyle: "bold" } }],
          [
            {
              content: "Required Investigations for Next Visit",
              colSpan: 6,
              styles: { fontStyle: "bold" },
            },
          ],
          ...investigationList,
          [
            { content: "Prepared By: " + header.CreatedUserName, colSpan: 3 },
            { content: "Referrals:", colSpan: 3 },
          ],
          [
            {
              content:
                (consultant.Name === null || consultant.Name === undefined
                  ? ""
                  : consultant.Name) +
                "\n" +
                (consultant.Qualifications === null ||
                consultant.Qualifications === undefined
                  ? ""
                  : consultant.Qualifications) +
                "\nConsultant\nSri Jayawardanapura General Hospital\nNugegoda, Sri Lanka",
              colSpan: 3,
              rowSpan: 2,
            },
            {
              content:
                "Next Clinic visit in: " +
                (header.NextClinicIn === null ||
                header.NextClinicIn === undefined ||
                header.NextClinicIn.includes("undefined")
                  ? ""
                  : header.NextClinicIn),
              colSpan: 3,
            },
          ],
          [
            {
              content:
                "Annual Visit: " +
                (header.AnnualVisit === null || header.AnnualVisit === undefined
                  ? ""
                  : header.AnnualVisit),
              colSpan: 3,
            },
          ],
        ],
      });
      doc.output("dataurlnewwindow");
    },
    async onShowReferralPDF() {
      let clinic = String(localStorage.getItem("Location"));
      const result = await this.api.get("/report/SummaryReport", {
        params: { DocID: String(localStorage.getItem("DDOCUMENT")) },
      });
      const result2 = await this.api.get("/username");
      let header = result.data.header;
      let consultant = result.data.consultant;
      let documentID = result.data.documentID;
      let referrals = result.data.refferalLetters;
      let referralList = [];
      for (let index = 0; index < referrals.length; index++) {
        const referral = referrals[index];
        let temp = referral.Letter.replaceAll('<br>', ' ').replaceAll('</p><p>', '/n').replaceAll('<p>', '').replaceAll('</p>', '').replaceAll('<br>', '').replaceAll('&nbsp;', ' ');
        var arr = temp.split('/n')
        for (let index = 0; index < arr.length; index++) {
          const contentText = arr[index];
          referralList.push([{ content: contentText, colSpan: 6 }]);
        }
        referralList.push([""]);
      }
      // A4 dimensions : 210 x 297
      const doc = new this.$PDF();
      doc.setProperties({
        title: "Referral Report",
        subject: "Report",
        author: "admin",
        keywords: "",
        creator: "SJGH",
      });

      let img = new Image();
      img.src = "/sjgh/logo.png";
      doc.addImage(img, "png", 15, 10, 20, 20);
      doc.autoTable({
        margin: { right: 0, left: 0 },
        styles: {
          fontStyle: "bold",
          fontSize: 16,
          cellPadding: 0.01,
        },
        columnStyles: {
          0: { halign: "center" },
        },
        theme: "plain",
        body: [
          [
            {
              content:
                clinic === "MC"
                  ? "Medical Clinic"
                  : "Endocrinology & Diabetes Centre",
              styles: { fontStyle: "bold", halign: "center" },
            },
          ],
          [{content:"Sri Jayewardenepura General Hospital",styles: { fontSize: 14, fontStyle: "normal", halign: "center" }}],
          [{content:"Thalapathpitiya, Nugegoda, Sri Lanka.",styles: { fontSize: 14,fontStyle: "normal",  halign: "center" }}],
          [
            {content:"Tel: 0112 778 610 Fax: 0112 778 611 Email: endocrine@sjghsrilanka.lk Web: sjgh.health.gov.lk	",styles: { fontSize: 12,fontStyle: "normal",  halign: "center" }},
          ],
          [""],
        ]
        });
      doc.autoTable({
        margin: { right: 2, left: 14 },
        styles: {
          fontStyle: "normal",
          cellPadding: 1,
          lineWidth: 0.01,
          lineColor: [255, 255, 255],
        },
        columnStyles: {
          0: { halign: "left", cellWidth: 35 },
          1: { halign: "left" },
          2: { halign: "left", cellWidth: 35 },
          3: { halign: "left", cellWidth: 35 },
          4: { halign: "left", cellWidth: 35 },
          5: { halign: "left" },
        },
        theme: "plain",
        body: [
          ...referralList,
          [
            {
              content:
                (consultant.Name === null || consultant.Name === undefined
                  ? ""
                  : consultant.Name) +
                "\n" +
                (consultant.Qualifications === null ||
                consultant.Qualifications === undefined
                  ? ""
                  : consultant.Qualifications) +
                "\nConsultant\nSri Jayawardanapura General Hospital\nNugegoda, Sri Lanka",
              colSpan: 3,
              rowSpan: 2,
            },
          ],
        ],
      });
      doc.output("dataurlnewwindow");
    },
    async GetClinicWiseElements() {
      let clinic = String(localStorage.getItem("Location"));
      const result = await this.api.get(
        `/Clinic-notes/GetClinicWiseElements?clinic=${clinic}`
      );
      let data = result.data.ClinicWiseElements;
      this.LabInvestInputData = []
      for (let i = 0; i < data.length; i++) {
        const element = data[i];
        let item = {
          code: element.cElementCode,
          label: element.vcElementName,
          unit: null,
          result: null,
          date: null
        };
        this.LabInvestInputData.push(item);
      }
    },
    async registrationWiseClinicSession() {
      let intial = { label: "New", value: "--" };

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
          label: x.DocumentNo.trim(),
          value: x.DocumentNo.trim(),
        }));
        this.ClinicSession.push(intial);
        this.ClinicWiseSession = this.ClinicSession.find(
          (x) => x.value.slice(6) === DateTime.now().toFormat("dd/MM/yy")
        );
        this.ClinicWiseSession =
          this.ClinicWiseSession === null ||
          this.ClinicWiseSession === undefined
            ? intial.value
            : this.ClinicSession.value;
      } else {
        this.ClinicSession = this.NewSession;
      }
      this.GetAutoloadFORdocID();
    },
    /*====================================================*/
    //                    SearchDoctor                    //
    /*====================================================*/
    SearchDoctor() {
      let nameSearchDr = this.ReferralForm.Doctor.toLowerCase();
      var data = this.GetDoctorNamelist.filter((data) => {
        return data.label.toLowerCase().includes(nameSearchDr);
      });
      this.DoctorNameFilter = data;
      //  console.log(this.DoctorValue,"Values")
      //  e.preventDefault()
      //  console.log(data,"HMHMH")
    },
    async LoadPrivilages() {
      const result = await this.api.get("/security/GetUserPrivilages");
      this.Privilages = result.data.Privilages;
      this.ShowProbAddOpt = this.Privilages.some(x => x.vcScreenCode == "DOCTOR" && x.vcEventCode == "2PROBADD");
    },
    // Done
    async GetDocumentHeaderProblemList() {
      let DocNo;
      DocNo = String(localStorage.getItem("DDOCUMENT"));
      //  GET DOCUMENT HEADER PROBLEM LIST
      const result = await this.api.get(
        `/Clinic-notes/GetDocumentHeaderProblemList?DocID=${DocNo}`
      );
      this.problemList = result.data.GetDocumentHeaderProblemList ?? [];
      for (let i = 0; i < this.problemList.length; i++) {
        this.problemList[i].Note = this.problemList[i].Note ?? "";
        if (
          this.problemList[i].ProblemCode === null &&
          this.problemList[i].ProblemDescription !== null
        ) {
          let val = this.diagnosticList.find(
            (x) =>
              x.value.Description.trim() ===
              this.problemList[i].ProblemDescription.trim()
          );
          if (!(val === null || val === undefined)) {
            this.problemList[i].ProblemCode = val.value.ID;
          }
        }
      }
      if (this.problemList.length === 0) {
        this.diagnosis.ProblemPriority = 1;
      } else {
        let priorities = this.problemList.map((x) => x.Priority);
        this.diagnosis.ProblemPriority = Math.max(...priorities) + 1;
      }
    },
    //ResetInputField
    ResetInput() {},
    async GetCodeOfDoct() {
      let t = this.ReferralForm.Doctor
      this.ReferralForm.DoctorID = this.DoctorNameFilter.find((data)=>{return(data.label === t)}).value
    },
    /*===========================================*/
    async GetDoctorName() {
      const result = await this.api.get(`/Clinic-notes/GetDoctorList`);
      let docNamelist = result.data.GetDoctorNamelist;
      if (docNamelist != null) {
        this.GetDoctorNamelist = docNamelist.map((data) => ({
          label: data.ConsultantName,
          value: data.ConsultantCode,
        }));
      }
      this.TodayDocumentNum();
    },

    async GetDocumentCheckedList() {
      let DocNo;
      DocNo = String(localStorage.getItem("DDOCUMENT"));
      //  GET DOCUMENT HEADER PROBLEM LIST
      const result = await this.api.get(
        `/Clinic-notes/GetDocumentCheckedList?DocID=${DocNo}`
      );
      this.checkedList = result.data.GetDocumentCheckedList;
      this.checkedList.sort(function (a, b) {
        if (a.ItemCode < b.ItemCode) {
          return -1;
        }
        if (a.ItemCode > b.ItemCode) {
          return 1;
        }
        return 0;
      });
      for (let index = 0; index < this.checkedList.length; index++) {
        let date = this.checkedList[index].Date;
        if (date !== null) {
          this.checkedList[index].value = true;
        } else {
          this.checkedList[index].value = false;
        }
      }
    },
    /*=============================*/
    //       DISEASE SAVE
    /*=============================*/
    async saveDiseaseofImage() {
      let form = {};
      form.doctorNoteID = String(localStorage.getItem("RegiNo"));
      form.nameOfDisease = this.disease.name;
      form.diseaseofNote = this.disease.discriptions;
      form.nameofImage = this.imagePreviewFordisease;
      form.userID = String(localStorage.getItem("UserId"));
      this.diseaseFinalSave = form;
      this.SaveClinicImageDisease();
      setTimeout(() => (this.imageModal = false), 1500);
      setTimeout(() => this.diseaseformReset(), 1500);
    },

    async GetDocumentPatientCheckedList() {
      let RegNo = String(localStorage.getItem("RegiNo"));
      let clinic = String(localStorage.getItem("Location"));
      //  GET DOCUMENT HEADER PROBLEM LIST
      const result = await this.api.get(
        `/Clinic-notes/GetDocumentPatientCheckedList?RegiNo=${RegNo}&Clinic=${clinic}`
      );
      this.checkedGrid = result.data.GetDocumentPatientCheckedList;
      if (this.checkedGrid !== null) {
        let keys = Object.keys(this.checkedGrid[0]);
        console.log(keys, "HUTEST");
        this.checkedGridFields = [];
        for (let i = 0; i < keys.length; i++) {
          const key = keys[i];
          let item = { filter: false, sorter: false, key: key, label: key };
          this.checkedGridFields.push(item);
        }
        for (let i = 0; i < this.checkedGrid.length; i++) {
          for (let t = 1; t < keys.length; t++) {
            let key = keys[t];
            this.checkedGrid[i][key] =
              this.checkedGrid[i][key] !== null ? "✔️" : "";
          }
        }
      }
    },
    async show() {
      let DocNo = String(localStorage.getItem("DDOCUMENT"));
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
      this.header.comment_print = BodyMI[0].Comment;
      this.header.comment_internal = BodyMI[0].CommentInternal;
      this.header.annualVisit = BodyMI[0].AnnualVisit;
      this.header.nextClinicIn = BodyMI[0].NextClinicIn;
      this.header.nextClinicInValue = this.header.nextClinicIn.split(" ")[0];
      this.header.nextClinicInUnit = this.header.nextClinicIn.split(" ")[1];
      await this.GetClinicMasterData();
      // this.BodyMI
    },
    async GetLabInvestigationData(global) {
      let RegiNo = localStorage.getItem("RegiNo");
      let ClinicCode = String(localStorage.getItem("Location"));
      let LabInvestfields = [
        {
          filter: false,
          sorter: false,
          key: "element_range",
          label: "Element Name / Range",
          _style: { width: "20%" },
        },
      ];

      let LabInvestData = [];
      {
        const result = await global.api.get(
          `/patient-biography/lab/LabResultDetail`,
          {
            params: {
              regcode: RegiNo,
              showall: global.lab_showall,
              clinic: ClinicCode,
            },
          }
        );
        if( result.data.LabResultDetail !== undefined){
        for (let i = 0; i < result.data.LabResultDetail.length; i++) {
          let x = result.data.LabResultDetail[i];
          let item = {};
          let element_id = x.cElementCode.trim();
          let element_range = `${x.vcElementName} | ${x.vcRange}`;
          if(this.lab_showall === true && this.lab_filter !== null){
            if(!element_range.includes(this.lab_filter)) continue;
          }
          item.element_id = element_id;
          item.element_range = element_range;
          item._cellClasses = {};
          let date = x.dtpCreateDate.split("T")[0].trim();
          let field = LabInvestfields.find(
            (t) => t.key === date + "_val"
          );
          if (field === null || field === undefined) {
            field = {
              filter: false,
              sorter: false,
              key: date + "_val",
              label: date,
            };
            LabInvestfields.push(field);
            LabInvestfields.sort((a, b) =>
              Date.parse(a.label) > Date.parse(b.label) ? -1 : 1
            );
          }
          let val = `${x.vcKeyword} ${x.vcUnit}`;
          item[date + "_val"] = val;
          item._cellClasses[date + "_val"] = "bg-success";
          let ref = LabInvestData.find(
            (t) => t.element_id === element_id
          );
          if (ref === null || ref === undefined) {
            LabInvestData.push(item);
          } else {
            ref._cellClasses[date + "_val"] = "bg-success";
            ref[date + "_val"] = val;
          }
        }
        }
      for (let i = 0; i < LabInvestData.length; i++) {
        let ref = LabInvestData[i];
        for (let i = 0; i < LabInvestfields.length; i++) {
          const key = LabInvestfields[i].key;
          if (ref[key] === undefined || ref[key] === "undefined") {
            ref[key] = "-";
          }
        }
      }
      }
      {
        const result = await global.api.get(
          `/patient-biography/lab/LabResultDetailExternal`,
          {
            params: {
              regcode: RegiNo,
              showall: global.lab_showall,
              clinic: ClinicCode,
            },
          }
        );
        if( result.data.LabResultDetail !== undefined){
        for (let i = 0; i < result.data.LabResultDetail.length; i++) {
          let x = result.data.LabResultDetail[i];
          let item = {};
          let element_id = x.cElementCode.trim();
          let element_range = x.vcElementName + " |\n" + x.vcRange;
          item.element_id = element_id;
          item.element_range = element_range;
          item._cellClasses = {};
          let date = x.dtpCreateDate.split("T")[0].trim();
          let field = LabInvestfields.find(
            (t) => t.key === date + "_val"
          );
          if (field === null || field === undefined) {
            field = {
              filter: false,
              sorter: false,
              key: date + "_val",
              label: date,
            };
            LabInvestfields.push(field);
            LabInvestfields.sort((a, b) =>
              Date.parse(a.label) > Date.parse(b.label) ? -1 : 1
            );
          }
          let val = x.vcKeyword + "\n" + x.vcUnit;
          item[date + "_val"] = val;
          item._cellClasses[date + "_val"] = "bg-primary";
          let ref = LabInvestData.find(
            (t) => t.element_id === element_id
          );
          if (ref === null || ref === undefined) {
            LabInvestData.push(item);
          } else {
            ref._cellClasses[date + "_val"] = "bg-primary";
            ref[date + "_val"] = val;
          }
        }
        }
      for (let i = 0; i < LabInvestData.length; i++) {
        let ref = LabInvestData[i];
        for (let i = 0; i < LabInvestfields.length; i++) {
          const key = LabInvestfields[i].key;
          if (ref[key] === undefined || ref[key] === "undefined") {
            ref[key] = "-";
          }
        }
      }
      }
      global.LabInvestfields = LabInvestfields
      global.LabInvestData = LabInvestData
    },
    // ==================================================
    async GetClinicMasterData() {
      let ClinicCode,
        DocumentID,
        DocID = String(localStorage.getItem("DocId"));
      DocumentID = String(localStorage.getItem("DDOCUMENT"));
      // DocumentID = "0039MC12/03/22      ";
      const resultForDosage = await this.api.get(`/Clinic-notes/GetDosageList`);
      this.GetDosageList = resultForDosage.data.GetDosageList.map((data) => ({
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
      for (let i = 0; i < this.DataforPrescriptons.length; i++) {
        let data = this.DataforPrescriptons[i];
        if (this.DataforPrescriptons[i].ItemType.trim() === "DRUG") {
          this.DataforPrescriptons[i].Strenth =
            data.Strenth === "NA" || data.Strenth === null
              ? null
              : data.Strenth.trim();
          this.DataforPrescriptons[i].Dosage =
            data.Dosage === "NA" || data.Dosage === null
              ? null
              : data.Dosage.trim();
          this.DataforPrescriptons[i].duration =
            data.duration === "NA" || data.duration === null
              ? null
              : data.duration.trim();
        } else {
          this.DataforPrescriptons[i].Strenth =
            data.Strenth === "NA" || data.Strenth === null
              ? "NA"
              : data.Strenth.trim();
          this.DataforPrescriptons[i].Dosage =
            data.Dosage === "NA" || data.Dosage === null
              ? "NA"
              : data.Dosage.trim();
          this.DataforPrescriptons[i].duration =
            data.duration === "NA" || data.duration === null
              ? "NA"
              : data.duration.trim();
        }
      }
      this.Tablelist();
      this.ReferralTemplates();
      this.ReferralForms();
      this.GetReferralbyLetterId();
    },
    /*================================================================*/
    // REFERRAL TEMPLATED
    /*================================================================*/
    async ReferralTemplates() {
      let ClinicCode = String(localStorage.getItem("Location"));
      const resultForReferralTemp = await this.api.get(
        `/Clinic-notes/GetDocumentReferralLetterTemplate?ClinicCode=${ClinicCode}`
      );
      var data = resultForReferralTemp.data.GetDocumentReferralLetterTemplate;
      if (data != null) {
        this.ReferralLetterTemplate =
          resultForReferralTemp.data.GetDocumentReferralLetterTemplate;
      }
    },
    /*================================================================*/
    // REFERRAL FORMS
    /*================================================================*/
    async ReferralForms() {
      let DocumentID = String(localStorage.getItem("DDOCUMENT"));
      const resultForReferralLetter = await this.api.get(
        `/Clinic-notes/GetDocumentReferralLetter?DocumentID=${DocumentID}`
      );
      var data = resultForReferralLetter.data.GetDocumentReferralLetter;
      if (data != null) {
        data.sort( (a,b) => Number(b.LetterID) - Number(a.LetterID) )
        this.ReferralLetter = data.map((data) => ({
          label: data.LetterID,
          value: data.LetterID,
        }));
        this.ReferralLetterform = data ;
        this.LetterID = this.ReferralLetterform[0].LetterID
        return (await this.GetReferralbyLetterId());
      }
    },

    /*================================================================*/
    //        PATIENT REGISTRATION INFO & PROFILE INFO
    /*================================================================*/
    async Promise() {
      let RegiNo = localStorage.getItem("RegiNo");
      const resultGetPatientRegistionInfo = await this.api.get(
        `/reception/GetPatientRegistrationInfo?RegiNo=${RegiNo}`
      );
      let data = resultGetPatientRegistionInfo.data.GetPatientRegistrationInfo;
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
        this.Profile = data;
      }
    },

    /*================================================================*/
    //        PATIENT DOCUMENT PROOF GET
    /*================================================================*/
    async GetDiseaseofImage() {
      let DocNo = this.$route.params.id;
      const resultImageOfDisease = await this.api.get(
        `Clinic-notes/GetPatientDiseaseOfImage?RegiNo=${DocNo}`
      );
      this.ImageOfDisease = resultImageOfDisease.data.GetPatientDiseaseofImage;
    },

    /*================================================================*/
    //             PATIENT PROFILE IMGAGE GET
    /*================================================================*/
    async GetProfile() {
      let RegiNo = localStorage.getItem("RegiNo");
      const result = await this.api.get(
        `Clinic-notes/GetPatientProfileImage?RegiNo=${RegiNo}`
      );
      let ProfileImage = result.data.GetPatientProfileImage;
      this.Photo = ProfileImage[0].Patient_Image;
    },

    /*================================================================*/
    //           GET PATIENT INITIAL SCREENING
    /*================================================================*/
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

    /*================================================================*/
    //           PROBLEM LIST  ADD
    /*================================================================*/
    addSpecialNote(diagnosis, problemList) {
      if (diagnosis.index !== null) {
        problemList[diagnosis.index].ProblemCode =
          diagnosis.ProblemDescriptionValue.ID;
        problemList[diagnosis.index].ProblemDescription =
          diagnosis.ProblemDescriptionValue.Description ?? "";
        problemList[diagnosis.index].Note = diagnosis.ProblemNote ?? "";
        problemList[diagnosis.index].Status = diagnosis.ProblemStatus
          ? "Active"
          : "Deactive";
        problemList[diagnosis.index].Priority = diagnosis.ProblemPriority;
      } else {
        let item = {};
        item.index = problemList.length;
        item.ProblemCode = diagnosis.ProblemDescriptionValue.ID;
        item.ProblemDescription =
          diagnosis.ProblemDescriptionValue.Description ?? "";
        item.Note = diagnosis.ProblemNote ?? "";
        item.Status = diagnosis.ProblemStatus ? "Active" : "Deactive";
        item.Priority = diagnosis.ProblemPriority;
        problemList.push(item);
      }
      diagnosis.ProblemNote = "";
      diagnosis.ProblemDescriptionLabel = null;
      diagnosis.ProblemDescriptionValue = null;
      diagnosis.ProblemStatus = true;
      diagnosis.ProblemPriority =
        Math.max(...problemList.map((o) => o.Priority), 0) + 1;
      diagnosis.index = null;
      this.SaveClinicSpecialNotes();
    },
    /*================================================================*/
    //           PROBLEM LIST  EDITE
    /*================================================================*/
    editSpecialNote(item, problemList) {
      this.diagnosis.index = problemList.indexOf(item);
      this.diagnosis.ProblemNote = item.Note ?? "";
      this.diagnosis.ProblemStatus =
        item.Status.trim() === "Active" ? true : false;
      this.diagnosis.ProblemPriority = item.Priority;
      this.diagnosis.ProblemDescriptionLabel = item.ProblemDescription.trim();
      this.GetDiagnosticList().then(() => {
        let val = this.diagnosticList.find(
          (x) => x.value.ID === Number(item.ProblemCode)
        );
        if (val === null || val === undefined) {
          val = this.diagnosticList.find(
            (x) => x.value.Description.trim() === item.ProblemDescription.trim()
          );
        }
        this.diagnosis.ProblemDescriptionValue = val.value;
        this.diagnosis.ProblemDescriptionLabel = val.label;
      });
    },
    /*================================================================*/
    //           PROBLEM LIST  DELETE
    /*================================================================*/
    deleteSpecialNote(item, problemList) {
      this.problemList = problemList.filter((a) => !(a === item));
      this.deleteDiagnosisItem = null;
      this.deleteDiagnosisModal = false;
      this.SaveClinicSpecialNotes();
    },

    // ad() {
    //   alert("hello");
    // },
    // deleteinfo() {
    //   alert("Delete function is unable fire");
    // },
    GetUsertypes() {
      let user = localStorage.getItem("Usertype");
      this.Usertype = user;
    },
    //*==============================================================================*/
    //                     GET VIEW BY SELECT CASE AND IMAGE                         //
    /*===============================================================================*/
    imageload(data) {
      var casedata = data.id;
      var finaldatacase = this.ImageOfDisease.filter(
        (data) => data.id === casedata
      );
      this.ImageOfDiseaseView = finaldatacase;
      this.warningModal = true;
    },
    //*==============================================================================*/
    //                         SAVE AS REFERRAL FORM                                 //
    /*===============================================================================*/
    async SaveasReferralform() {
      let Inden,
        ReferralLetterform = this.ReferralLetterform;
      if (ReferralLetterform != null) {
        const ids = ReferralLetterform.map((object) => {
          return object.LetterID;
        });
        const Formid = Math.max(...ids);
        Inden = Formid + 1;
      } else {
        Inden = 1;
      }
      var form = {};
      form.documentID = String(localStorage.getItem("TodayDoc"));
      form.letterID = Inden;
      form.doctorID = this.ReferralForm.DoctorID;
      form.letter = this.content;
      this.PatientReferralbyID = form;

      await this.SaveClinicReferral();
      // After save Reset Field
      this.ReferralReset();
    },
    //*==============================================================================*/
    //                         UPDATE AS REFERRAL FORM                                 //
    /*===============================================================================*/
    async UpdateasReferralform() {
      var form = {};
      form.documentID = String(localStorage.getItem("TodayDoc"));
      form.letterID = this.LetterID;
      form.doctorID = this.ReferralForm.DoctorID;
      form.letter = this.content;
      this.PatientReferralbyID = form;

      await this.SaveClinicReferral();
      // After save Reset Field
      this.ReferralReset();
    },
    /*===============================================================================*/
    //                          SAVE AS TEMPLATE                                    //
    /*===============================================================================*/
    /*===============================================================================*/
    //                          SAVE AS TEMPLATE                                    //
    /*===============================================================================*/
    async SaveasTemplate() {
      let form = {};
      form.templateName = this.ReferralTemp;
      form.doctorID = this.ReferralForm.DoctorID;
      form.letter = this.content;
      form.clinicID = localStorage.getItem("Location");
      this.ReferralLetterTemplateArray = form;

      setTimeout(() => (this.ReferraltemplateName = false), 1500);

      form = this.ReferralLetterTemplateArray;
      this.api
        .post("Clinic-notes/SaveClinicReferralTemplate", form)
        .then((result) => {
          if (result.data === "success") {
            this.alertboxSuccess = true;
            setTimeout(() => (this.alertboxSuccess = false), 3000);
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
      // After save Reset Field
      // this.ReferralReset();
      // This function call back to internal reload
      this.ReferralTemplates();
    },

    /* ==== Call referral template modal ==== */
    async CallRefNameTempModel() {
      this.ReferraltemplateName = true;
    },
    //==========

    loadRefTemplate(item) {
      this.ReferralForm.Letter = item.Letter;
      setTimeout((this.loadtemplate = false), 1500);
    },

    /*=======================================================*/
    //           DISEASE UPLOAD FORM RESET
    /*=======================================================*/
    diseaseformReset() {
      this.imagePreviewFordisease = null;
      this.disease.name = "";
      this.disease.discriptions = "";
    },
    /*=======================================================*/
    //          BMI FORM RESET
    /*=======================================================*/
    healthNBMI() {
      this.header.bp = null;
      (this.header.pulse_rate = null), (this.header.weight = 0);
      this.header.height = 0;
      this.header.bmi = 0;
      this.header.classificationValue = null;
      this.header.status = null;
      this.header.status_com = null;
      this.header.comment_print = null;
      this.header.comment_internal = null;
      this.header.nextClinicIn = null;
      this.header.nextClinicInValue = null;
      this.header.nextClinicInUnit = null;
      this.header.annualVisit = null;
    },
    /*=======================================================*/
    //          REFERRAL RESET
    /*=======================================================*/
    ReferralReset() {
      this.ReferralForm.Doctor = "";
      this.ReferralForm.Letter = "";
      this.content = "";
      this.ReferralLetter = null;
    },
    /*=======================================================*/
    //          INITIAL SCREEN RESET
    /*=======================================================*/
    initialformreset() {
      this.initScreen.id = 0;
      this.initScreen.presentcomplaint = null;
      this.initScreen.past_medi_or_surgicalHx = null;
      this.initScreen.obstretic_or_menstrualHx = null;
      this.initScreen.medication_summary = null;
      this.initScreen.familtyHistory = null;
      this.initScreen.allergicHx = null;
      this.initScreen.socialHx = null;
      this.initScreen.generalExamination = null;
      this.initScreen.smorking = null;
      this.initScreen.alcohol = null;
      this.initScreen.fGS = null;
      this.initScreen.fGS_Image = null;
      this.initScreen.Thyroid_Examination = null;
      this.initScreen.eyes = null;
      this.initScreen.rS = null;
      this.initScreen.aBD = null;
      this.initScreen.cNS = null;
      this.initScreen.cVS = null;
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
      this.header.bmi =
        this.header.weight /
        ((this.header.height / 100) * (this.header.height / 100));
      this.header.bmi = Math.round(this.header.bmi * 100) / 100;
    },

    check(item) {
      const val = Boolean(this.usersData[item.id]._selected);
      this.$set(this.usersData[item.id], "_selected", !val);
      console.log(item, "DataforInvest");
    },

    closeIm() {
      this.imageModal = false;
    },

    loadTemplate() {
      this.loadtemplate = true;
      this.ReferralTemplates();
    },

    async onDiagnosisSelect() {
      let val = this.diagnosticList.find(
        (x) =>
          x.value.Description.trim() ===
          this.diagnosis.ProblemDescriptionLabel.trim()
      );
      this.diagnosis.ProblemDescriptionValue = val.value;
    },
    /*=====================================================*/
    //   GET Referral by LetterId
    /*=====================================================*/
    async GetReferralbyLetterId() {
      let Get1,
        form = this.ReferralLetterform;
      Get1 = this.LetterID;
      let data = form.filter((data) => data.LetterID === Get1);
      var finalData = data[0].Letter;
      this.ReferralForm.DoctorID = data[0].DoctorID;
      return (this.content = finalData);
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
      this.iternalReload();
      this.show();
      if (DocumentNo === "--") {
        localStorage.setItem("DDOCUMENT", DocumentNo);

        // this.healthNBMI()
        // this.iternalReload()
        setTimeout(() => this.show(), 2000);
      } else {
        let route_params = this.$route.params.id;
        this.ClinicWiseSession = "--";
        let route = this.$router.resolve({
          path: `/doctor-queue-or-portal/${route_params}/Notes_`,
        });
        window.open(route.href, "_blank");
        localStorage.setItem("DDOCUMENT", DocumentNo);
      }
      // this.iternalReload()
      // this.show()
    },

    /*===================================================*/
    //             NEW Record save and
    /*===================================================*/
    async GetAutoloadFORdocID() {
      let regi, clinic;
      regi = String(localStorage.getItem("RegiNo"));
      clinic = String(localStorage.getItem("Location"));
      const result = await this.api.get(
        `/Clinic-notes/LastWiseClinicSessions?regno=${regi}&clinic=${clinic}`
      );
      let Data,
        LastDOCNum = result.data.RegistrationWiseClinicSessions;
      // if(LastDOCNum != null){
      Data = LastDOCNum[0].DocumentID.trim();
      localStorage.setItem("DDOCUMENT", Data);
      this.DOC_ID = localStorage.getItem("DDOCUMENT");
      // }
    },

    showDocNo() {
      this.ClinicWiseSession = localStorage.getItem("DDOCUMENT", DocumentNo);
    },

    /*================================================================*/
    //           HERE FUNCTION CALL FOR INTERNAL-RELOAD
    /*================================================================*/

    async iternalReload() {
      this.GetDocumentHeaderProblemList();
      this.GetClinicMasterData();
    },

    async onDrugSelect() {
      if ("Drugs" === this.Prescript.drugsSearch) {
        this.DrugsPropertyAutocomplete();
      }else{
        let text = this.Prescript.InputSearch;
        if(text.includes(" | Profile")){
          text = text.replace(" | Profile","");
        }
        if ("Lab" === this.Prescript.drugsSearch) {
          const result = await this.api.get(
            `/Clinic-notes/GetFilterLabInvestlist`,
            { params: { SearchText: text } }
          );
          const filtered = result.data.GetFilterLablist.find(x=>x.vcItemDescrition.trim() === text.trim());
          this.Prescript.selectedDrug = filtered !== null || filtered !== undefined ? filtered : null;
      } else {
          const result = await this.api.get(
            `/Clinic-notes/GetFilterOtherlist`,
            { params: { SearchText: text } }
          );
          const filtered = result.data.GetFilterOtherInvestigationlist.find(x=>x.vcItemDescrition.trim() === text.trim());
          this.Prescript.selectedDrug = filtered !== null || filtered !== undefined ? filtered : null;
        }
      }
    },

    /*================================================================*/
    //   MULTI SELECT SEARCH [ DRUGS || LAB || OTHER] FOR PRESCRIPTION
    /*================================================================*/

    async onDrugSearch(e) {
      let val = this.Prescript.InputSearch;
      let Stext = val.length;
      let i = 0;
      this.InputSearchSelect = null;
      if ("Drugs" === this.Prescript.drugsSearch) {
        i = 4;
        if (i < Stext) {
          let text = val;
          const result = await this.api.get(`/Clinic-notes/GetFilterDruglist`, {
            params: { SearchText: text },
          });
          const filterd = result.data.GetFilterDrugsList;
          let mo = filterd.map((data) => {
            let item = data;
            item.value = data.ATCClessification.trim();
            item.label =
              data.cItemCode.trim() +
              " " +
              data.ATCClessification.trim() +
              " " +
              data.vcItemDescrition.trim();
            item.itemCode = data.cItemCode.trim();
            item.Form = data.Form;
            return item;
          });
          return (this.InputSearchSelect = mo !== undefined ? mo  : this.InputSearchSelect);
        }
      } else if ("Lab" === this.Prescript.drugsSearch) {
        i = 2;
        if (i < Stext) {
          let text = val;
          const result = await this.api.get(
            `/Clinic-notes/GetFilterLabInvestlist`,
            { params: { SearchText: text } }
          );
          const filterd = result.data.GetFilterLablist;
          let mo = filterd.map((data) => {
            let item = data;
            item.value = data.vcItemDescrition.trim();
            item.label =
              data.cItemCode.trim() + " " + data.vcItemDescrition.trim();
            item.itemCode = data.cItemCode.trim();
            return item;
          });
          return (this.InputSearchSelect = mo !== undefined ? mo  : this.InputSearchSelect);
        }
      } else if ("Other" === this.Prescript.drugsSearch) {
        i = 2;
        if (i < Stext) {
          let text = val;
          const result = await this.api.get(
            `/Clinic-notes/GetFilterOtherlist`,
            { params: { SearchText: text } }
          );
          const filterd = result.data.GetFilterOtherInvestigationlist;
          let mo = filterd.map((data) => {
            let item = data;
            item.value = data.vcItemDescrition.trim();
            item.label =
              data.cItemCode.trim() + " " + data.vcItemDescrition.trim();
            item.itemCode = data.cItemCode.trim();
            return item;
          });
          return (this.InputSearchSelect = mo !== undefined ? mo  : this.InputSearchSelect);
        }
      }
    },

    /*================================================================*/
    //   IF SELECT LAB AND OTHER; THEN DISABLE TO UNNECESSARY ELEMENT
    /*================================================================*/
    DisabledForPrescriptionOtherField() {
      if ("Lab" === this.Prescript.drugsSearch) {
        document.getElementById("Strenth").disabled = true;
        document.getElementById("Unit").disabled = true;
        document.getElementById("Frequncy").disabled = true;
        document.getElementById("Duration").disabled = true;
        document.getElementById("Urgent").disabled = false;
      } else if ("Other" === this.Prescript.drugsSearch) {
        document.getElementById("Strenth").disabled = true;
        document.getElementById("Unit").disabled = true;
        document.getElementById("Frequncy").disabled = true;
        document.getElementById("Duration").disabled = true;
        document.getElementById("Urgent").disabled = false;
      } else if ("Drugs" === this.Prescript.drugsSearch) {
        document.getElementById("Strenth").disabled = false;
        document.getElementById("Unit").disabled = false;
        document.getElementById("Frequncy").disabled = false;
        document.getElementById("Duration").disabled = false;
        document.getElementById("Urgent").disabled = false;
      }
    },

    /*================================================================*/
    //               DRUG PROPERTY AUTO COMPLETE
    /*================================================================*/
    async DrugsPropertyAutocomplete() {
      const INPUTTEXT =
        this.Prescript.InputSearch ??
        this.Prescript.InputSearch;
      const result = await this.api.get(`/Clinic-notes/GetDrugAutoComplete`, {
        params: { GenericName: INPUTTEXT },
      });
      const drugautoComplete = result.data.DrugAutoCompleteI;
      this.Prescript.selectedDrug = drugautoComplete[0];
      if(this.Prescript.selectedDrug.Cream === true){
        document.getElementById("Strenth").disabled = true;
        document.getElementById("Unit").disabled = true;
        document.getElementById("Frequncy").disabled = true;
        document.getElementById("Duration").disabled = true;
        document.getElementById("Urgent").disabled = false;
      }
      let unit = this.Prescript.selectedDrug.Unit;
      if(this.Prescript.index === null){
      this.Prescript.strenth = (
        this.Prescript.selectedDrug.Strength ?? ""
      ).trim();
      this.Prescript.selUnit = (unit ?? "").trim();
      }
    },

    /*================================================================*/
    //             OBJECT TRIM  || PRESCRIPTION ADD , EDITE AND DELETE
    /*================================================================*/
    AddForPrescription() {
      let type = this.Prescript.drugsSearch.trim();
      let allowAdd = true;
      if (
        type === "Drugs" &&
        (this.Prescript.InputSearch === null ||
          this.Prescript.InputSearch.trim() === "" ||
          this.Prescript.selFreque === null ||
          this.Prescript.week === null ||
          this.Prescript.strenth === null ||
          this.Prescript.selUnit === null)
      ) {
        this.AlertMessage = "Please select drug and required fields";
        this.alertModal = true;
        allowAdd = false;
      }
      if (
        type === "Drugs" &&
        (this.Prescript.InputSearch === null ||
          this.Prescript.InputSearch.trim() === "")
      ) {
        this.AlertMessage = "Please select drug and required fields";
        this.alertModal = true;
        allowAdd = false;
      }
      if (
        type !== "Drugs" &&
        type !== null &&
        (this.Prescript.InputSearch === null ||
          this.Prescript.InputSearch.trim() === "")
      ) {
        this.AlertMessage = "Please select investigation";
        this.alertModal = true;
        allowAdd = false;
      }
      if (
        type === "Drugs" && this.Prescript.index === null && (this.Prescript.selectedDrug === null )
      ) {
        this.AlertMessage = "Please Select Proper drug from List";
        this.alertModal = true;
        allowAdd = false;
      }
      if (
        (type === "Lab" && this.Prescript.index === null) && (this.Prescript.selectedDrug === null )
      ) {
        this.AlertMessage = "Please Select Proper Lab Investigation from List";
        this.alertModal = true;
        allowAdd = false;
      }
      if (
        (type === "Other" && this.Prescript.index === null)  && (this.Prescript.selectedDrug === null )
      ) {
        this.AlertMessage = "Please Select Proper Other Investigation from List";
        this.alertModal = true;
        allowAdd = false;
      }
      if (allowAdd) {
        switch (type) {
          case "Drugs":
            type = "DRUG";
            break;
          case "Lab":
            type = "LAB";
            break;
          case "Other":
            type = "OTHER";
            break;
          default:
            break;
        }
        if (this.Prescript.index === null) {
          let item = {};
          item.documentID = String(localStorage.getItem("DDOCUMENT"));
          item.ID = 0;
          item.comment = this.Prescript.comment;
          item.Description = this.Prescript.InputSearch;
          item.ItemType = type.trim();
          if (item.ItemType.trim() === "DRUG") {
            item.ItemCode =
            this.Prescript.selectedDrug == null
              ? this.Prescript.InputSearch.split(" ")[0] ??
                this.Prescript.InputSearch
              : this.Prescript.selectedDrug.GenericId;
            item.Dosage = this.Prescript.selFreque;
            item.duration = this.Prescript.week;
            item.Strenth =
              this.Prescript.strenth + " " + this.Prescript.selUnit;
          } else {
            item.ItemCode =
            this.Prescript.selectedDrug == null
              ? this.Prescript.InputSearch.split(" ")[0] ??
                this.Prescript.InputSearch
              : this.Prescript.selectedDrug.cItemCode;
            item.Dosage = "NA";
            item.duration = "NA";
            item.Strenth = "NA";
          }
          item.Urgent = this.Prescript.urgentCheck ?? false;
          item.Sequence =
            this.DataforPrescriptons === null
              ? 1
              : this.DataforPrescriptons.length + 1;
          this.DataforPrescriptons.push(item);
        } else {
          let index = this.Prescript.index;
          this.DataforPrescriptons[index].DocumentID = String(
            localStorage.getItem("DDOCUMENT")
          );
          this.DataforPrescriptons[index].ItemCode =
            this.Prescript.selectedDrug == null
              ? this.Prescript.InputSearch.split(" ")[0] ??
                this.Prescript.InputSearch
              : this.Prescript.selectedDrug.GenericId;
          this.DataforPrescriptons[index].ItemType = type.trim();
          this.DataforPrescriptons[index].comment = this.Prescript.comment;
          this.DataforPrescriptons[index].Description =
            this.Prescript.InputSearch;
          if (this.DataforPrescriptons[index].ItemType.trim() === "DRUG") {
            this.DataforPrescriptons[index].Dosage = this.Prescript.selFreque;
            this.DataforPrescriptons[index].duration = this.Prescript.week;
            this.DataforPrescriptons[index].Strenth =
              this.Prescript.strenth + " " + this.Prescript.selUnit;
          } else {
            this.DataforPrescriptons[index].Dosage = "NA";
            this.DataforPrescriptons[index].duration = "NA";
            this.DataforPrescriptons[index].Strenth = "NA";
          }
          this.DataforPrescriptons[index].Urgent =
            this.Prescript.urgentCheck ?? false;
          this.DataforPrescriptons[index].Sequence =
            this.DataforPrescriptons === null
              ? 1
              : this.DataforPrescriptons.length + 1;
        }
        this.InputSearchSelect = null;
        this.Prescript.index = null;
        this.Prescript.selectedDrug = null;
        this.Prescript.selFreque = null;
        this.Prescript.InputSearch = null;
        this.Prescript.comment = null;
        this.Prescript.selUnit = null;
        this.Prescript.strenth = null;
        this.Prescript.urgentCheck = false;
        if (type.trim() !== "DRUG") {
          this.Prescript.week = null;
        }
        this.DisabledForPrescriptionOtherField();
        this.Tablelist();
        this.SaveClinicPrescription();
      }
    },

    /*================================================================*/
    //             OBJECT TRIM  || PRESCRIPTION EDITE !
    /*================================================================*/
    editPrescription(vitem) {
      let item = this.DataforPrescriptons[vitem.idx];
      this.Prescript.index = this.DataforPrescriptons.indexOf(item);
      this.Prescript.id = item.ID;
      let type = item.ItemType.trim();
      switch (type) {
        case "DRUG":
          this.Prescript.drugsSearch = "Drugs";
          break;
        case "LAB":
          this.Prescript.drugsSearch = "Lab";
          break;
        case "OTHER":
          this.Prescript.drugsSearch = "Other";
          break;
        default:
          break;
      }
      this.Prescript.InputSearch = item.Description;
      this.onDrugSelect();
      this.Prescript.comment = item.comment;
      this.Prescript.urgentCheck = item.Urgent;
      if (type === "DRUG") {
        this.Prescript.selFreque = item.Dosage.trim() ?? "";
        this.Prescript.week = item.duration;
        this.Prescript.strenth = item.Strenth.split(" ")[0];
        this.Prescript.selUnit = item.Strenth.split(" ")[1];
      } else {
        this.Prescript.selFreque = null;
        this.Prescript.week = null;
        this.Prescript.strenth = null;
        this.Prescript.selUnit = null;
      }
      this.DisabledForPrescriptionOtherField();
    },

    /*================================================================*/
    //             PRESCRIPTION DELETE !
    /*================================================================*/
    async deletePrescription(vitem) {
      let item = this.DataforPrescriptons[vitem.idx];
      this.DataforPrescriptons = this.DataforPrescriptons.filter(
        (a) => !(a === item)
      );
      this.deletePrescriptionItem = null;
      this.deletePrescriptionModal = false;
      this.DisabledForPrescriptionOtherField();
      this.Tablelist();
      this.SaveClinicPrescription();
    },
    async TodayDocumentNum() {
      var registrationNo = localStorage.getItem("RegiNo");
      var clinicCode = localStorage.getItem("Location");
      const result = await this.api.get(
        `Clinic-notes/GetTodayDocumentNumber?RegiNo=${registrationNo}&ClinicCode=${clinicCode}`
      );
      let ToDayDocument = result.data.TodayDocument;
      localStorage.setItem("TodayDoc", ToDayDocument[0].DocumentNo);
    },

    /*================================================================*/
    //             PRESCRIPTION ITEMS FILTER BY URGENT & NORMAL
    /*================================================================*/
    async Tablelist() {
      /* PRESCRIPTIONS */
      this.DataforPrescriptonslist = this.DataforPrescriptons.filter(
        (data) => data.Urgent !== true
      ).map((data) => ({
        Dosage:
          (data.Strenth ?? "NA") +
          " | " +
          (data.Dosage ?? "NA") +
          " | " +
          (data.duration ?? "NA"),
        Comment: data.comment,
        Description: data.Description,
        ItemType: data.ItemType,
        ID: data.ID,
        idx: this.DataforPrescriptons.indexOf(data),
      }));

      /* URGENT PRESCRIPTIONS */
      this.UrgentPrescriptionDatalist = this.DataforPrescriptons.filter(
        (data) => data.Urgent === true
      ).map((data) => ({
        Dosage:
          (data.Strenth ?? "NA") +
          " | " +
          (data.Dosage ?? "NA") +
          " | " +
          (data.duration ?? "NA"),
        Comment: data.comment,
        Description: data.Description,
        ItemType: data.ItemType,
        ID: data.ID,
        idx: this.DataforPrescriptons.indexOf(data),
      }));
    },

    /*================================================================*/
    //            GET GENERAL DIAGNOSTIC LIST
    /*================================================================*/
    async GetDiagnosticList() {
      let text = this.diagnosis.ProblemDescriptionLabel ?? "";
      if (text.length > 2) {
        const result = await this.api.get(`Clinic-notes/GetDiagnosticList`, {
          params: { SearchText: text },
        });
        this.diagnosticList = result.data.GetDiagnosticList.map((data) => ({
          label: data.Description,
          value: data,
        }));
      } else {
        this.diagnosticList = [];
      }
    },

    /*==============================================*/
    //           FINAL TAB WISE SAVE                //
    /*==============================================*/
    async onSaveClick() {
      await this.SaveClinicHeader();
      // await this.SaveClinicImageDisease();
      {
        await this.SaveClinicInitScreening();
      }
      {
        await this.SaveClinicSpecialNotes();
      }
      {
        await this.SaveClinicCheckedList();
      }
      {
        await this.SaveClinicPrescription();
      }
      {
        // await this.SaveClinicReferralTemplate();
        // await this.SaveClinicReferral();
        // await this.ReferralReset();
      }
    },
    /*================================================================*/
    //       FINAL POST DATA TO BACKEND-ENDPOINT [CLINIC HEADER]
    /*================================================================*/
    async SaveClinicHeader() {
      if (
        this.header.nextClinicInUnit === undefined ||
        this.header.nextClinicInValue === undefined ||
        this.header.nextClinicInUnit === null ||
        this.header.nextClinicInValue === null
      ) {
        this.header.nextClinicIn = null;
      } else {
        this.header.nextClinicIn = `${this.header.nextClinicInValue} ${this.header.nextClinicInUnit}`;
      }
      let form = {};
      form.registrationNo = String(localStorage.getItem("RegiNo"));
      form.documentID = String(localStorage.getItem("TodayDoc"));
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
            this.registrationWiseClinicSession();
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

    /*================================================================*/
    //  FINAL POST DATA TO BACKEND-ENDPOINT [CLINIC INITIAL SCREENING]
    /*================================================================*/
    async SaveClinicInitScreening() {
      let form = {};
      form.id = this.initScreen.id;
      form.presentingComplaint = this.initScreen.presentcomplaint ?? "";
      form.pastMedicalHx = this.initScreen.past_medi_or_surgicalHx ?? "";
      form.obstreticHx = this.initScreen.obstretic_or_menstrualH ?? "";
      form.MedicationSummary = this.initScreen.medication_summary ?? "";
      form.familyHistory = "";
      form.allegicHx = this.initScreen.allergicHx ?? "";
      form.socialHx = this.initScreen.socialHx ?? "";
      form.occupation = "";
      form.smoking = this.initScreen.smorking ?? "";
      form.alcohol = this.initScreen.alcohol ?? "";
      form.generalExamination = this.initScreen.generalExamination ?? "";
      form.FGS = this.initScreen.fGS ?? "";
      form.FGSImage = this.initScreen.fGS_Image ?? "";
      form.ThyroidExam = this.initScreen.Thyroid_Examination ?? "";
      form.Eyes = this.initScreen.eyes ?? "";
      form.RS = this.initScreen.rs ?? "";
      form.ABD = this.initScreen.aBD ?? "";
      form.CNS = this.initScreen.cNS ?? "";
      form.CVS = this.initScreen.cVS ?? "";
      form.clinicType = String(localStorage.getItem("Location"));
      form.registrationID = String(localStorage.getItem("RegiNo"));
      form.userID = String(localStorage.getItem("UserId"));
      this.api
        .post("Clinic-notes/SaveClinicInitScreening", form)
        .then((result) => {
          if (result.data === "success") {
            this.alertboxSuccess = true;
            setTimeout(() => (this.alertboxSuccess = false), 3000);
            this.registrationWiseClinicSession();
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
      this.SaveClinicCheckedList();
    },

    /*================================================================*/
    //  FINAL POST DATA TO BACKEND-ENDPOINT [CHECK-LIST]
    /*================================================================*/
    async SaveClinicCheckedList() {
      let form = {};
      form.documentID = String(localStorage.getItem("TodayDoc"));
      let items = this.checkedList.filter((x) => x.value === true);
      form.itemIDs = items.map((x) => ({
        itemID: x.ItemCode.trim(),
      }));
      this.api
        .post("Clinic-notes/SaveClinicCheckList", form)
        .then((result) => {
          if (result.data === "success") {
            this.alertboxSuccess = true;
            setTimeout(() => (this.alertboxSuccess = false), 3000);
            this.show();
            this.GetDocumentCheckedList();
            this.GetDocumentPatientCheckedList();
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

    /*================================================================*/
    //  FINAL POST DATA TO BACKEND-ENDPOINT [CHECK-LIST]
    /*================================================================*/
    async SaveClinicExternalLabData() {
      let form = {};
      form.registrationID = String(localStorage.getItem("RegiNo"));
      form.userID = String(localStorage.getItem("UserId"));
      let items = this.LabInvestInputData.filter((x) => x.result !== null );
      form.elements = items.map((x) => ({
        elementCode:x.code,
        result:x.result,
        date:x.date,
        unit:x.unit
      }));
      this.api
        .post("Clinic-notes/SaveClinicExternalLabData", form)
        .then((result) => {
          if (result.data === "success") {
            this.alertboxSuccess = true;
            setTimeout(() => (this.alertboxSuccess = false), 3000);
            this.show();
            this.GetLabInvestigationData(this);
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
    /*================================================================*/
    //  FINAL POST DATA TO BACKEND-ENDPOINT [PRESCRIPTION]
    /*================================================================*/
    async SaveClinicPrescription() {
      let form = {};
      form.documentID = String(localStorage.getItem("TodayDoc"));
      form.prescriptions = this.DataforPrescriptons;
      for (let i = 0; i < form.prescriptions.length; i++) {
        form.prescriptions[i].Sequence = i;
      }
      this.api
        .post("Clinic-notes/SaveClinicPrescription", form)
        .then((result) => {
          if (result.data === "success") {
            this.alertboxSuccess = true;
            setTimeout(() => (this.alertboxSuccess = false), 3000);
            this.GetAutoloadFORdocID();
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

    /*================================================================*/
    //  FINAL POST DATA TO BACKEND-ENDPOINT [SPECIALNOTES]
    /*================================================================*/
    async SaveClinicSpecialNotes() {
      let form = {};
      form.registrationID = String(localStorage.getItem("RegiNo"));
      form.documentID = String(localStorage.getItem("TodayDoc"));
      form.specialnotes = this.problemList;
      this.api
        .post("Clinic-notes/SaveClinicSpecialNotes", form)
        .then((result) => {
          if (result.data === "success") {
            this.alertboxSuccess = true;
            setTimeout(() => (this.alertboxSuccess = false), 3000);
            this.GetAutoloadFORdocID();
            this.registrationWiseClinicSession();
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
      let form = this.PatientReferralbyID;
      this.api
        .post("Clinic-notes/SaveClinicReferral", form)
        .then((result) => {
          if (result.data === "success") {
            this.alertboxSuccess = true;
            setTimeout(() => (this.alertboxSuccess = false), 3000);
            this.show();
            this.GetAutoloadFORdocID();
            this.registrationWiseClinicSession();
            this.ReferralTemplates();
            this.ReferralForms();
            this.GetReferralbyLetterId();
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
      let form;
      form = this.ReferralLetterTemplateArray;
      this.api
        .post("Clinic-notes/SaveClinicReferralTemplate", form)
        .then((result) => {
          if (result.data === "success") {
            this.alertboxSuccess = true;
            setTimeout(() => (this.alertboxSuccess = false), 3000);
            this.show();
            this.registrationWiseClinicSession();
            this.SaveasTemplate();
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
    //                  FINAL SAVE FOR DISEASE
    /*====================================================================*/

    async SaveClinicImageDisease() {
      let form = this.diseaseFinalSave;
      this.api
        .post("Clinic-notes/SaveClinicImageDisease", form)
        .then((result) => {
          if (result.data === "success") {
            this.alertboxSuccess = true;
            setTimeout(() => (this.alertboxSuccess = false), 3000);
            this.GetDiseaseofImage();
            this.registrationWiseClinicSession();
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
    async AddNewProblem() {
      let form = this.NewProblem;
      this.api
        .post("Clinic-notes/AddNewProblem", form).then((result) => {
          if (result.data === "success") {
            this.alertboxSuccess = true;
            setTimeout(() => (this.alertboxSuccess = false), 3000);
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
    }
  },
  created() {
    this.LoadPrivilages();
    this.TodayDocumentNum();
    this.registrationWiseClinicSession();
    this.GetClinicWiseElements();
    this.GetLabInvestigationData(this);
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
    this.GetDoctorName();
    this.showDocNo();
  },
};
</script>


<style></style>
