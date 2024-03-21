<template>
  <div>
   
    <CContainer fluid >
      <CRow>
        <CCol style="width: 100%; padding: 0">
          <CCard style="height:100%"> 
            <CCardBody >
              <CDataTable
                header
                hover
                striped
                :items="appointments"
                :fields="fields"
                :items-per-page="15"
                clickable-rows
                @row-clicked="call"
                :pagination="{ doubleArrows: false, align: 'center' }"
                column-filter
                table-filter
                sorter
                items-per-page-select
              >
                <template #call="{ item }">
                  <!-- d{{item.id}}  -->
                  <td class="py-10">
                    <CButtonGroup>
                      <div size="sm" class="mt-6" @click="call(item)">
                        <CButton
                          type="button"
                          color="success"
                          size="sm"
                          class="font-weight-bold1"
                          >Call</CButton
                        >
                      </div>
                    </CButtonGroup>
                  </td>
                </template>

                <template #Comment="{item}">
                 <!-- <div>
                  
                 </div> -->
                  <td class="py-10">
                     <p>{{item.Comment}}</p>
                    <p v-if="item.ModifyUser" class="text-primary" style="font-size:11px;">Last Seen : {{item.ModifyUser}}</p>
                    <p v-else></p>
                  </td>
                </template>
              </CDataTable>
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
          label: "Registration No",
        },
        { filter: false, sorter: false, key: "PatientName", label: "Name" },
        { filter: false, sorter: false, key: "SerAppNo", label: "Appointment No" },
        { filter: false, sorter: false, key: "Service", label: "Service" },
        { filter: false, sorter: false, key: "ClinicAppNo", label: "Token No" },
        { filter: false, sorter: false, key: "Priority", label: "Priority" },
        { filter: false, sorter: false, key: "Comment", label: "Comment" },
        // { filter: false, sorter: false, key: "Priority", label: "Priority" },
        {
          key: "call",
          label: "",
          _style: { width: "1%" },
          filter: false,
          sorter: false,
          _props: { color: "primary", className: "fw-semibold" },
        },
      ],
      appointments: [],
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
  computed: {
   
  },

  methods: {
    sortBy: function (sortKey) {
      this.reverse = this.sortKey == sortKey ? !this.reverse : false;
      this.sortKey = sortKey;
    },
    cancel() {
      this.patient = null;
    },

    
    
   call(item) {
     setTimeout( () => this.$router.push({path:`/doctor-queue-or-portal/${item.RegistrationNo}`}), 1);
      // let DOC = item.DocumentNo.replaceAll("/","%2F")
       let DOC = item.DocumentNo
      localStorage.setItem("DDOCUMENT",DOC)
      localStorage.setItem("RegiNo",item.RegistrationNo)

    },

    async getAppointments() {
      let Ccode = localStorage.getItem("Location");
      const result = await this.api.get(`doctorportal/appointments?clinic=${Ccode}`);
      this.appointments = result.data.appointments;
    },

    
  },

  created() {
    this.getAppointments();
    
  },
};
</script>