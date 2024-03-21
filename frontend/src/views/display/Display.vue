<template>
  <div class="row wallpaper dev" style="height: 100%; padding: 0">
    <div class="col align-self-center text-center" style="width: 100%">
      <h2 class="text-left strokeme" style="font-size: 24pt;width:100%;color:white;">
        {{this.clinicName}}<br>{{this.date}}
      </h2>
      <video muted preload="auto" ref="videoplayer" style="width:100%" @ended="videoEnded">
        <source ref="videosrc">
      </video>
    </div>
    <div class="col-2 " style="height: 100%; padding: 6px;width:100%">
      <h2 class="text-center  strokeme" style="font-size: 24pt;width:100%;color:white">
        NEW
      </h2>
      <CCard color="blue" style="margin:2px" v-bind:key="appointment.vcPatientName" v-for="appointment in appointmentsNew">
          <CCardHeader align="center" style="font-size: 300%;font-weight:bold;padding:0px;margin:2px">
            {{Number(appointment.ClinicAppNo).toString().padStart(2, '0')}}
          </CCardHeader>
        </CCard>
    </div>
    <div class="col-2 " style="height: 100%; padding: 6px;width:100%">
      <h2 class="text-center  strokeme" style="font-size: 24pt;width:100%;color:white">
        FOLLOW UP
      </h2>
      <CCard color="blue" style="margin:2px" v-bind:key="appointment.vcPatientName" v-for="appointment in appointmentsFollowUP">
          <CCardHeader align="center" style="font-size: 300%;font-weight:bold;padding:0px;margin:2px">
            {{Number(appointment.ClinicAppNo).toString().padStart(2, '0')}}
          </CCardHeader>
        </CCard>
    </div>
  </div>
</template>

<script>

export default {
  name: "Display",
  inject: ["api"],
  data: function () {
    return {
      active: null,
      room_no: null,
      break: 1,
      channels: [],
      clinicName: null,
      clinic: "D&E",
      appointmentsFollowUP: [],
      appointmentsNew: [],
      videoindex: 0,
      beforeValue: 0,
      PresentValue: 0,
      a: [],
      b: [],
      date: (new Date()).toLocaleString('en-GB'),
      fields: [
        { filter: false, sorter: false, key: "token_no", label: "Token" ,width:"10%"},
        {
          filter: false,
          sorter: false,
          key: "doctor_name", width:"10%",
          label: "Doctor Name",
        }, 
        { filter: false, sorter: false, key: "room_no", label: "Room" },  
      ],
      videofiles: [],
      audiofiles: [],
    };
  },
 
  updated(){
  },

  created() {
    this.clinic = this.$route.query.clinic
    this.onStartup();
      setInterval(async () => {
        this.date = (new Date()).toLocaleString('en-GB')
      }, 100);
			setInterval(async () => {
        await this.getAppointments();
      }, 5000);
  },
  methods: {
    videoEnded() {
      this.videoindex++;
      if(this.videoindex >= this.videofiles.length){
        this.videoindex = 0;
      }
      this.$refs.videosrc.src = this.videofiles[this.videoindex].src;
      this.$refs.videosrc.type = this.videofiles[this.videoindex].type;
      this.$refs.videoplayer.load()
      this.$refs.videoplayer.play()
    },
    async getAppointments() {
      let result = await this.api.get("display/appointments", { params: { clinic: this.clinic } });
      let appointments = result.data.appointments;
      this.appointmentsNew = appointments.filter(x=> x.cSessionFor.trim() === "NEW") ?? []
      this.appointmentsFollowUP = appointments.filter(x=> x.cSessionFor.trim() === "FOLLOW-UP") ?? []
      this.appointmentsNew.sort(function(a, b) {
         return a.ClinicAppNo - b.ClinicAppNo;
      });
      this.appointmentsFollowUP.sort(function(a, b) {
         return a.ClinicAppNo - b.ClinicAppNo;
      });
    },
    async loadVideos() {
      this.videofiles = [];
      let result = await this.api.get("display/videolist");
      let files = result.data;
      let baseurl = this.api.defaults.baseURL
      files.forEach((file) => {
        let video = {
          id: file,
          src: baseurl + "/display/videos/" + file,
          type: "video/mp4",
        };
        this.videofiles.push(video);
      });
    },

    async onStartup() {
      const result = await this.api.get(
        `user/locations?userid=hbt`
    );
      this.locations = result.data;
      let location = this.locations.find((data) => data.code.trim() === this.clinic);
      this.clinicName = location.label
      await this.loadVideos();
      this.$refs.videosrc.src = this.videofiles[this.videoindex].src;
      this.$refs.videosrc.type = this.videofiles[this.videoindex].type;
      this.$refs.videoplayer.load()
      this.$refs.videoplayer.play()
      await this.getAppointments();
    },
  },
};
</script>

<style>
.strokeme
{
    color: white;
    text-shadow:
    -1px -1px 0 #000,
    1px -1px 0 #000,
    -1px 1px 0 #000,
    1px 1px 0 #000;  
}
</style>
<style lang="sass" scoped>
.wallpaper
  position: relative
  background: #979797
  overflow: hidden
.wallpaper .q-page-container
  position: relative
  z-index: 2
  overflow: hidden
.wallpaper:before
  content: ' '
  display: block
  position: absolute
  width: 100%
  height: 100%
  z-index: 0
  opacity: 29
.wallpaper.dev:before
 
  background: url(~../../assets/image/back6.jpg) repeat
</style>