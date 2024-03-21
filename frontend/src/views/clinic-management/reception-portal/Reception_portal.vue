<template>
  <div>
    <CToaster position="top-center">
      <CToast :show="alertboxSuccess" title="CToast static component">
        Successful Submit
      </CToast>
      <CToast :show="alertboxFailed" title="CToast static component">
        Failed Submit - {{err}}
      </CToast>
    </CToaster>
    <CContainer fluid style="height: 100%">
      <CRow>
        <!-- ko{{this.appointments}} -->
        <CCol style="width: 100%; padding: 0">
          <CCard style="height: 100%">
            <CCardBody>
              <!-- Helo{{appointments}} -->
              <CDataTable
                header
                clickableRows
                :tableProps="{
                  striped: true,
                  hover: true,
                }"
                hover
                striped
                :items="appointments"
                :fields="fields"
                :items-per-page="15"
                @row-clicked="show"
                :pagination="{ doubleArrows: false, align: 'center' }"
                column-filter
                table-filter
                sorter
                items-per-page-select
              >
                <template #show_details="{ item }">
                  <!-- d{{item.id}}  | clickable-rows -->
                  <td class="py-10">
                    <CButtonGroup>
                      <div size="sm" class="mt-6" @click="show(item)">
                        <CButton
                          type="button"
                          color="info"
                          size="sm"
                          class="font-weight-bold1"
                          >Accept</CButton
                        >
                      </div>
                    </CButtonGroup>
                  </td>
                </template>
              </CDataTable>
            </CCardBody>
          </CCard>
        </CCol>

        <CCol
          v-if="patient !== null"
          style="width: 100%; margin-left: 3px; margin-right: -2px; padding: 0"
        >
          <CCard style="height: 100%">
            <CCardBody>
              
              <CRow>
                <CCol>
                  <CInput
                    label="Patient Name"
                    v-model="patient.PatientName"
                    readonly
                  />
                </CCol>
              </CRow>
              <CRow>
                <CCol>
                  <CInput
                    label="Category"
                    v-model="patient.Category"
                    readonly
                  />
                </CCol>
                <CCol>
                  <CInput
                    label="Registration No"
                    v-model="patient.RegistrationNo"
                    readonly
                  />
                </CCol>
              </CRow>
              <CRow>
                <CCol>
                  <CInput label="Age" v-model="patient.Age" readonly />
                </CCol>
                <CCol>
                  <CInput label="Gender" v-model="patient.Gender" readonly />
                </CCol>
                <CCol>
                  <CInput
                    label="Reference No"
                    v-model="patient.ReferenceNo"
                    readonly
                  />
                </CCol>
              </CRow>
              <CRow>
                <CCol>
                  <CInput
                    label="Appointment No"
                    v-model="patient.AppNo"
                    readonly
                  />
                </CCol>
                <CCol>
                  <CInput
                    label="Appointment Date"
                    v-model="patient.AppDate"
                    readonly
                  />
                </CCol>
              </CRow>
              <CRow>
                <CCol class="d-flex">
                  <div class="mr-2">Priority Levels</div>
                  &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
                  <CInputRadioGroup
                    class="ml-10"
                    inline
                    :options="priorities"
                    :checked.sync="patient.Priority"
                    style="margin-left: 40px"
                  />
                </CCol>
              </CRow>
              <CRow>
                <CCol> </CCol>
              </CRow>
              <CRow>
                <CCol>
                  <CInput label="Comment" v-model="patient.Comment" />
                </CCol>
              </CRow>
              <CRow>
                <CCol style="margin-right:-250px;">
                  <CButton @click="this.submit" style="width: 7.5rem" color="behance">
                    Save
                  </CButton>
                </CCol>
                <CCol>
                  <CButton @click="this.cancel" style="width: 7.5rem" color="reddit" align="center">
                    Cancel
                  </CButton>
                </CCol>
              </CRow>
            </CCardBody>
          </CCard>
        </CCol>
      </CRow>
    </CContainer>
  </div>
</template>



<script>
export default {
  name: "ReceptionPortal",
  inject: ["api"],

  data() {
    return {
      fields: [
        {
          filter: false,
          sorter: false,
          key: "RegistrationNo",
          label: "Reg No",
        },
        { filter: false, sorter: false, key: "PatientName", label: "Name" },
        { filter: false, sorter: false, key: "Gender", label: "Gender" },
        { filter: false, sorter: false, key: "Age", label: "Age" },
        { filter: false, sorter: false, key: "AppNo", label: "App No" },
        // { filter: false, sorter: false, key: "Priority", label: "Priority" },
        { filter: false, sorter: false, key: "SessionFor", label: "Session" },
        {
          key: "show_details",
          label: "",
          _style: { width: "1%" },
          filter: false,
          sorter: false,
          _props: { color: "primary", className: "fw-semibold" },
        },
      ],
      appointments: [],
      patient: null,
      priorities: [
        { value: "regular", label: "Regular" },
        { value: "priority", label: "Priority" },
        { value: "immediate", label: "Immediate" },
      ],
      err: '',
      alertboxSuccess: false,
      alertboxFailed: false,
      session: [],
      id: null,
      sortKey: "",
      search: "",
      reverse: false,
      SessionList: [
        { label: "FOLLOW-UP" },
        { label: "FIRST-VISIT" },
        { label: "STAFF" },
      ],

      errors: [],
    };
  },

  // -------------------COMPUTED----------------------
  computed: {},

  methods: {
    sortBy: function (sortKey) {
      this.reverse = this.sortKey == sortKey ? !this.reverse : false;
      this.sortKey = sortKey;
    },
    cancel() {
      this.patient = null;
    },
    submit() {
      let form = {}
      form.referenceNo = this.patient.ReferenceNo.trim();
      form.pr_level = this.patient.Priority.trim();
      form.comment = this.patient.Comment.trim();
      form.user = localStorage.getItem("UserId");
      this.api
        .post("reception/patient", form)
        .then((result) => {
          console.log(result.status);
          if (result.data === "success") {
            this.alertboxSuccess = true;
            setTimeout(() => (this.alertboxSuccess = false), 3000);
            this.patient = null;
            this.getAppointments();
          } else {
            this.alertboxFailed = true;
            this.err = "Could not save patient Reference No." + form.referenceNo
            setTimeout(() => (this.alertboxFailed = false), 3000);
          }
        })
        .catch((error) => {
          this.err = error;
          this.alertboxFailed = true;
          setTimeout(() => (this.alertboxFailed = false), 3000);
        });
    },
    async show(item) {
      var today = new Date();
      let date =
        today.getFullYear() +
        "-" +
        (today.getMonth() + 1) +
        "-" +
        today.getDate();
        // let date =
      let Ccode = localStorage.getItem("Location");
      // let result = await this.api.get(`reception/patient/${item.AppNo}/${date}}/${Ccode}/${item.ReferenceNo}`);
      let result = await this.api.get(`reception/patients?AppNo=${item.AppNo}&ClinicDate=${date}&clinic=${Ccode}&ReferenceNo=${item.ReferenceNo}`);
      let patient = result.data;
      if (patient !== null) {
        switch (patient.Gender) {
          case "M":
            patient.Gender = "Male";
            break;
          case "F":
            patient.Gender = "Female";
            break;

          default:
            break;
        }
        patient.Priority = this.priorities[patient.Priority].value;
        patient.AppDate = new Date(patient.AppDate).toLocaleDateString("en-GB");
        this.patient = patient;
      }
    },

    async getAppointments() {
      var today = new Date();
      let date =
        today.getFullYear() +
        "-" +
        (today.getMonth() + 1) +
        "-" +
        today.getDate();
      let Ccode = localStorage.getItem("Location");
      const result = await this.api.get(`reception/GetClinicappointments?clinic=${Ccode}&date=${date}`);
      this.appointments = result.data.appointments;
      // alert("ooo");
    },
  },

  created() {
    this.getAppointments();
    // alert("");
  },
  rowClicked() {
    console.log("event-click");
  },
};
</script>
