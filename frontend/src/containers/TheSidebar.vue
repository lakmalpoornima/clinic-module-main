<template>
  <CSidebar
    :minimize="minimize"
    unfoldable
    :show="show"
    @update:show="(value) => $store.commit('set', ['sidebarShow', value])"
  >
  
    <CSidebarBrand class="d-md-down-none" to="/">
       
       <CHeaderNavLink>
        <div class="c-avatar c-sidebar-brand-full" size="custom-size"
        :height="25" viewBox="0 0 642 134">
          <img src="img/avatars/user.png" class="c-avatar-img" />
       
        </div>

        <span class="c-sidebar-brand-full text-center text-uppercase font-weight-bold text-white p-0 mb" size="custom-size"
       
        viewBox="0 0 642 134" style="margin-top:10px; padding:0px;" >{{CurrentUser}}</span>
      </CHeaderNavLink>

     
    </CSidebarBrand>
    <CRenderFunction flat :contentToRender="sidebarItems"/>
  </CSidebar>
</template>

<script>
export default {
  name: 'TheSidebar',
  inject: ["api"],

  data () {
    return {
      CurrentUser: null,
      Privilages: null
      }
  },

  computed: {
    show () {
      return this.$store.state.sidebarShow
    },
    minimize () {
      return this.$store.state.sidebarMinimize
    },
    sidebarItems () {
      let items = []
      if (this.Privilages != null)
      {
        let ShowDashOpt = this.Privilages.some(x => x.vcScreenCode == "DASH" && x.vcEventCode == "1VIEW");
        let ShowDocOpt = this.Privilages.some(x => x.vcScreenCode == "DOCTOR" && x.vcEventCode == "1VIEW");
        let ShowUserOpt = this.Privilages.some(x => x.vcScreenCode == "USER" && x.vcEventCode == "1VIEW");
        let ShowBioOpt = this.Privilages.some(x => x.vcScreenCode == "PATBIO" && x.vcEventCode == "1VIEW");
        let ShowRecepOpt = this.Privilages.some(x => x.vcScreenCode == "RECEP" && x.vcEventCode == "1VIEW");
        items =  [
        {
          _name: "CSidebarNav",
          _children: [
            ShowDashOpt ? {
              _name: "CSidebarNavItem",
              name: this.$t("menu.dashboard"),
              to: "/dashboard",
              icon: "cil-applications-settings",
            } : {},
             {
              _name: "CSidebarNavItem",
              name: "Display",
              to: `/display?clinic=${encodeURIComponent(String(localStorage.getItem("Location")))}`,
              icon: "cil-notes"
              },
            // -------------------------------------------------------------------
            {
              _name: "CSidebarNavTitle",
              _children: ["Clinic Management"],
            },

            ShowRecepOpt ? {
              _name: "CSidebarNavItem",
              name: "Reception Portal",
              to: "/Reception-Portal",
              icon: "cil-user-female",
              items:[
                 {
                  name: "Patient biography",
                  to: "/doctor-queue-or-portal/:id",
                },
              ]
            } : {},
            ShowDocOpt ? {
              _name: "CSidebarNavItem",
              name: "Doctor Queue / Portal",
              to: "/doctor-queue-or-portal",
              icon: "cil-people",
            } : {},

            ShowBioOpt ? {
              _name: "CSidebarNavItem",
              name: "Patient biography",
              to: "/doctor-queue-or-portal/:id",
              icon: "cil-notes",
            } : {},
            ShowDocOpt ? {
              _name: "CSidebarNavDropdown",
              name: " Clinic Visit / Notes",
              route: "/Clinic Management",
              icon: "cil-cursor",
              items: [
                // {
                //   name: "Initial Screening",
                //   to: "/Clinic Management/clinic-visit-or-notes/:id/Initial-Screening",
                // },

                {
                  name: "Prescription / Notes",
                  // to: "/Clinic Management/doctor-queue-or-portal/:id/Prescription",
                },
                {
                  name: "Prescription /Notes",
                  to: "/Clinic Management/doctor-queue-or-portal/:id/Prescription_",
                },
                // {
                //   name: "Comments and Image",
                //   to: "/Clinic Management/comments and image",
                // },
                // {
                //   name: "Check list",
                //   to: "/Clinic Management/doctor-queue-or-portal/:id/Check list",
                // },
                // {
                //   name: "Pregnancy",
                //   to: "/Clinic Management/doctor-queue-or-portal/:id/Pregnancy",
                // },
              ],
            } : {},
            // End  Clinic Management
            // -------------------------------------------------------------------
          ],
        },
      ];
      }
      return items 
    }
  },
  
  methods:{
    async LoadUser() {
      const result = await this.api.get("/username");
      this.CurrentUser = result.data;
    },
    async LoadPrivilages() {
      const result = await this.api.get("/security/GetUserPrivilages");
      this.Privilages = result.data.Privilages;
    },
  },

  created() {
    this.LoadUser();
    this.LoadPrivilages();
  },

}
</script>

