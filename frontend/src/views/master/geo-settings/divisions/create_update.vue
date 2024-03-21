<template>
  <div>
  
    <CRow>

      
      <CCol sm="12" md="6">
        <CCard>
          <CCardHeader>
             <div  v-if='this.id'>
            <strong>Update</strong><small>Form</small>
            </div>
            <div v-else>
              <strong>Division</strong> Form
            </div>
            
          </CCardHeader>

          
           <CCardBody>
            <CInput placeholder="Company Code" 
                valid-feedback="Input is valid."
                invalid-feedback="Please provide a Company Code."
                value=""
                :is-valid="validator">
              <template #prepend-content><CIcon name="cil-user"/></template>
            </CInput>

            <CInput placeholder="Company Name" 
                valid-feedback="Input is valid."
                invalid-feedback="Please provide Company Name."
                value=""
                :is-valid="validator">
              <template #prepend-content><CIcon name="cil-user"/></template>
            </CInput>

            <CInput placeholder="Address" 
                valid-feedback="Input is valid."
                invalid-feedback="Please provide a an Address."
                value=""
                :is-valid="validator">
              <template #prepend-content><CIcon name="cil-user"/></template>
            </CInput>

            <CInput placeholder="Telephone" 
                valid-feedback="Input is valid."
                invalid-feedback="Please provide Telephone."
                value=""
                :is-valid="validator">
              <template #prepend-content><CIcon name="cil-user"/></template>
            </CInput>

            <CInput placeholder="E-mail" 
                valid-feedback="Input is valid."
                invalid-feedback="Please provide a E-mail."
                value=""
                :is-valid="validator">
              <template #prepend-content><CIcon name="cil-user"/></template>
            </CInput>

            <CInput placeholder="Fax" 
                valid-feedback="Input is valid."
                invalid-feedback="Please provide a Fax."
                value=""
                :is-valid="validator">
              <template #prepend-content><CIcon name="cil-user"/></template>
            </CInput>

            <CInput placeholder="URL" 
                valid-feedback="Input is valid."
                invalid-feedback="Please provide a Url."
                value=""
                :is-valid="validator">
              <template #prepend-content><CIcon name="cil-user"/></template>
            </CInput>

            
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
            
          
        </CCard>
      </CCol>
    </CRow>

    


    
    
    
    
  </div>
</template>

<script>
export default {
  name: 'Forms',
  data () {
    return {
      selected: [], // Must be an array reference!
      show: true,
      horizontal: { label:'col-3', input:'col-9' },
      options: ['Option 1', 'Option 2', 'Option 3'],
      selectOptions: [
        'Option 1', 'Option 2', 'Option 3',
        { 
          value: 'some value', 
          label: 'Selected option'
        }
      ],
      selectedOption: 'some value',

      formCollapsed: true,
      checkboxNames: ['Checkboxes', 'Inline Checkboxes',
                      'Checkboxes - custom', 'Inline Checkboxes - custom'],
      radioNames: ['Radios', 'Inline Radios',
                    'Radios - custom', 'Inline Radios - custom']
    }
  },
  methods: {
    validator (val) {
      return val ? val.length >= 4 : false
    },
    back:function(){
      //  console.log(getlocal,'IDD')
       this.$router.push('/dashboard')
    }
  }

  
}
</script>
