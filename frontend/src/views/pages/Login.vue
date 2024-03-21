<template v-for="(color, key) in ['success', 'primary', 'info', 'warning', 'danger']">
  <div
    class="c-app flex-row align-items-center container"
    :class="{ 'c-dark-theme': $store.state.darkMode }"
  >
    <CToaster position="top-center">
      <CToast
        :show="alertboxSuccess"
        color="info"
        title="CToast static component"
      >
        Successful Login
      </CToast>
      <CToast
        :show="alertboxFailed"
        color="info"
        title="CToast static component"
      >
        Failed Login
      </CToast>
    </CToaster>
    <CContainer style="width: 40%">
      <CForm>
        <div class="row">
          <h4 class="col align-self-center text-center">
            SysTolic Clinic Module
          </h4>
        </div>
        <div class="row">
          <CInput
            v-model="username"
            placeholder="Username"
            type="text"
            class="col align-self-center"
            @input="this.onUserNameInput"
          >
            <template #prepend-content><CIcon name="cil-user" /> </template>
          </CInput>
        </div>
        <div class="row">
          <CInput
            class="col align-self-center"
            v-model="password"
            placeholder="Password"
            type="password"
          >
            <template #prepend-content>
              <CIcon name="cil-lock-locked" />
            </template>
          </CInput>
        </div>
        <div class="row">
          <CSelect
            class="col align-self-center"
            :value.sync="location"
            :options="locations"
            @click="getServiceLocation"
          >
          </CSelect>
        </div>
        <div class="row">
          <CButton class="col align-self-center" color="info" @click="login">
            Login
          </CButton>
        </div>
      </CForm>
    </CContainer>
  </div>
</template>

<script>
export default {
  name: "Login",
  inject: ["api"],

  data: function () {
    return {
      locations: null,
      location: null,
      username: null,
      err: "",
      password: null,
      alertboxFailed: false,
      alertboxSuccess: false,
    };
  },
  computed: {
    LocationIsvalid() {
      return !!this.location;
    },
    UsernameIsvalid() {
      return !!this.username;
    },
    PasswordIsvalid() {
      return !!this.password;
    },
    FormSubmitLogin() {
      return (
        this.UsernameIsvalid && this.LocationIsvalid && this.PasswordIsvalid
      );
    },
  },

  methods: {
    async login() {
      if (this.FormSubmitLogin) {
        let result = await this.api.post("/login", {
          Userid: this.username,
          Password: this.password,
          Location: this.location,
        });
        if (result.data === "success") {
          localStorage.setItem("UserId", this.username.trim());
          localStorage.setItem("Location", this.location.trim());
          this.getServiceLocation();
          this.alertboxSuccess = true;
          setTimeout(() => (this.alertboxSuccess = false), 3000);
          this.$router.push("/dashboard");
        } else {
          this.alertboxFailed = true;
          setTimeout(() => (this.alertboxFailed = false), 3000);
        }
      }
    },

    async getLocation() {
      const result = await this.api.get(
        `user/locations?userid=${this.username}`
      );
      this.locations = result.data;
      if (this.locations != null) {
        this.locations = this.locations.map((x) => {
          let container = {};
          container.label = x.label;
          container.value = x.code;
          return container;
        });
        this.location = this.locations[0].value;
      }
    },
    /* Find Serive location */
    async getServiceLocation() {
      let data = this.locations;
      let sortdata = data.filter((data) => data.value === this.location);
      let serve = sortdata[0].label;
      return localStorage.setItem("Servelocation", serve);
    },

    onUserNameInput() {
      this.getLocation();
    },
    close() {
      setTimeout((this.loginFailed = false), 1000);
    },
  },
};
</script>
