import Vue from 'vue'
import Router from 'vue-router'


const TheContainer = () =>
    import ('@/containers/TheContainer')

// Views
const Dashboard = () =>
    import ('@/views/Dashboard')
const Widgets = () =>
    import ('@/views/widgets/Widgets')
const Display = () =>
    import ('@/views/display/Display')
// Views - Pages
const Page404 = () =>
    import ('@/views/pages/Page404')
const Page500 = () =>
    import ('@/views/pages/Page500')
const Login = () =>
    import ('@/views/pages/Login')
const Register = () =>
    import ('@/views/pages/Register')

// Users
const Users = () =>
    import ('@/views/users/Users')
const User = () =>
    import ('@/views/users/User')


// for create and edit form of user
const UserAction = () =>
    import ('@/views/users/Create');

// -------------------------------------------------------------------------------
// --------------MASTER DATA GEO SETTING COMPONENT WITH ROUTES---------START------
// -------------------------------------------------------------------------------

const Company_List = () =>
    import ('@/views/master/geo-settings/company-info/CompanyList')

const Company_info_create_n_update = () =>
    import ('@/views/master/geo-settings/company-info/Company-info')

const Departments_List = () =>
    import ('@/views/master/geo-settings/departments/Departments-list')

const Create_or_Update_Department = () =>
    import ('@/views/master/geo-settings/departments/Create_or_update')


const Areas_List = () =>
    import ('@/views/master/geo-settings/areas/create_update')

// -------------------------------------------------------------------------------
// --------------MASTER DATA GEO SETTING COMPONENT WITH ROUTES---------END------
// -------------------------------------------------------------------------------


// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++



// -------------------------------------------------------------------------------
// --------------MASTER DATA  SETTING COMPONENT WITH ROUTES--------START------
// -------------------------------------------------------------------------------

const Initial_Screening = () =>
    import ('@/views/clinic-management/clinic_visit_or_notes/Initial_Screening/Initial_screening')

const Model = () =>
    import ('@/views/clinic-management/clinic_visit_or_notes/Initial_Screening/Model')

const Doctor_portal = () =>
    import ('@/views/clinic-management/doctor-queue-or-portal/Doctor_portal')

const Patient_biography = () =>
    import ('@/views/clinic-management/patient-biography/Patient_biography')

const Reception_Portal = () =>
    import ('@/views/clinic-management/reception-portal/Reception_portal')

const Comment_image = () =>
    import ('@/views/clinic-management/clinic_visit_or_notes/Comments and Image/comment_image')

const Check_list = () =>
    import ('@/views/clinic-management/clinic_visit_or_notes/Check list/Check_list')

const Pregnancy = () =>
    import ('@/views/clinic-management/clinic_visit_or_notes/Pregnancy/Pregnancy')

// const Notes = () =>
//     import ('@/views/clinic-management/clinic_visit_or_notes/Prescription Notes/Notes.vue')

const Prescription_notes = () =>
    import ('@/views/clinic-management/clinic_visit_or_notes/Prescription Notes/Prescription_notes')

const ModelPN = () =>
    import ('@/views/clinic-management/clinic_visit_or_notes/Prescription Notes/PrescriptionNotes')
    // const Refferals = () =>

const Lab_investication = () =>
    import ('@/views/clinic-management/lab_investigation/Laboratory_investigation')

// FOR HIDE CRUD BUTTON.....
const Prescription_notes_hide = () =>
    import ('@/views/clinic-management/clinic_visit_or_notes/Prescription Notes/Prescription_notes_hide')




// -------------------------------------------------------------------------------
// --------------MASTER DATA GEO SETTING COMPONENT WITH ROUTES---------END------Lab_investigation
// -------------------------------------------------------------------------------

Vue.use(Router)

export default new Router({
        mode: 'hash',
        // hash
        // https://router.vuejs.org/api/#mode
        linkActiveClass: 'open active',
        scrollBehavior: () => ({ y: 0 }),
    routes: [{
        path: '/',
        redirect: 'login',
    },
    {
        path: '/login',
        name: 'Login',
        component: Login
    },
    {
        path: '/display',
        name: 'Display',
        component: Display
    },
    {
        path: '/register',
        name: 'Register',
        component: Register
    },
    {
        path: '/404',
        name: '404',
        component: Page404
    },
    {
        path: '/500',
        name: '500',
        component: Page500
    },
    // This for System is Navigation

    {
        path: '/',
        // name: 'Home',
        component: TheContainer,
        children: [{
            path: 'dashboard',
            name: 'Dashboard',
            component: Dashboard
        },]
    },


    {
        path: '/system-security',
        name: 'System Security',
        component: TheContainer,
        children: [{
            path: 'users',
            meta: { label: 'UsersE' },
            component: {
                render(c) { return c('router-view') }
            },
            children: [{
                path: '',
                name: 'Users',
                component: Users
            },

            {
                path: '/system-security/UserAction/123',
                name: 'UserAction',
                component: UserAction
            },

            {
                path: 'User/:id',
                meta: {
                    label: 'User Edit'
                },
                name: 'UserAction',
                component: UserAction,
                props: true
            },

            {
                path: '/system-security/User/:id',
                name: 'User',
                component: UserAction,
                props: true
            },

            ]
        },]
    },


    {
        path: '/geo-settings',
        name: 'Geo Settings',
        component: TheContainer,
        children: [{
            path: 'company',
            meta: { label: 'Company' },
            component: {
                render(c) { return c('router-view') }
            },

            children: [{
                path: '/geo-settings/company info',
                name: 'Company',
                component: Company_List
            }, {
                path: '/geo-settings/company',
                name: 'Company',
                component: Company_info_create_n_update
            }, {
                path: '/geo-settings/company/:id',
                name: 'Company',
                component: Company_info_create_n_update
            }]
        },
                    
        {
            path: 'deprtments',
            meta: { label: 'Departments' },
            component: {
                render(c) { return c('router-view') }
            },

            children: [{
                path: '/geo-settings/departments',
                name: 'Departments',
                component: Departments_List
            },
            // if need create new record [  POST ]
            {
                path: '/geo-settings/department',
                name: 'Departments',
                component: Create_or_Update_Department
            },
            // update a single record [ PUT ]
            {
                path: '/geo-settings/department/:id',
                name: 'Departments',
                component: Create_or_Update_Department
            }

            ]

        },

        {
            path: 'areas',
            meta: { label: 'Areas' },
            component: {
                render(c) { return c('router-view') }
            },

            children: [{
                path: '/geo-settings/areas',
                name: 'Areas',
                component: Areas_List
            }]
        }
        ]
    },

    {
        path: '/',
        component: TheContainer,
        children: [
            {
                path: '/Reception-Portal',
                name: 'Reception Portal',
                component: Reception_Portal,
            }
        ]
    },
    
    {
        path: '/',
        component: TheContainer,
        children: [
            {
                path: '/doctor-queue-or-portal',
                name: 'Doctor Portal',
                component: Doctor_portal,
            },
            {
                path: '/doctor-queue-or-portal/:id',
                name: 'Patient biography',
                component: Patient_biography,

            },
            {
                path: '/doctor-queue-or-portal/:id/Notes',
                name: 'Prescription Notes',
                component: Prescription_notes

            },
            // FOR HIDE THE CRUD BUTTONS
            {
                path: '/doctor-queue-or-portal/:id/Notes_',
                name: 'Patient biography [ View Only ]',
                component: Prescription_notes_hide,
            },
        ]
        },
    
            {
                path: '/Clinic Management',
                name: 'Clinic Management',
                component: TheContainer,

                children: [{
                        path: 'company',
                        // meta: { label: 'Company1' },
                        component: {
                            render(c) { return c('router-view') }
                        },

                        children: [{
                                path: '/Clinic Management/clinic-visit-or-notes/:id/Initial-Screening',
                                name: 'Initial Screening',
                                component: Model,
                            },
                            {
                                path: '/Clinic Management/comments and image',
                                name: '',
                                component: Comment_image,

                            },
                            {
                                path: '/Clinic Management/doctor-queue-or-portal/:id/Check list',
                                name: 'Reception-Portal',
                                component: Check_list,

                            },

                            {
                                path: '/Clinic Management/doctor-queue-or-portal/:id/Pregnancy',
                                name: '',
                                component: Pregnancy

                            },

                            {
                                path: '/Clinic Management/doctor-queue-or-portal/:id/Notes',
                                name: '',
                                component: Prescription_notes

                            },

                            // {
                            //     path: '/Clinic Management/doctor-queue-or-portal/:id/Notes_',
                            //     name: '',
                            //     component: Lab_investication
                            // },
                            

                            {
                                path: '/Clinic Management/doctor-queue-or-portal/:id/Prescription',
                                name: '',
                                component: ModelPN

                            },
                            // {
                            //     path: '/Clinic Management/Patient_biography/:id',
                            //     name: '',
                            //     component: Notes

                            // },

                            {
                                path: '/Clinic Management/doctor-queue-or-portal/:id/lab-investigations',
                                name: 'lab-investigation',
                                component: Lab_investication

                            },
                            {
                                path: '/Clinic Management/doctor-queue-or-portal/:id/drugs_admin',
                                name: 'Drugs Admin',
                                component: Lab_investication
                            },
                            {
                                path: '/Clinic Management/doctor-queue-or-portal/:id/radiology',
                                name: 'Radiology',
                                component: Lab_investication
                            },
                            {
                                path: '/Clinic Management/doctor-queue-or-portal/:id/clinic_admissions',
                                name: 'Clinic Admissions',
                                component: Lab_investication
                            },




                        ]
                    },

                ]
            },

            // CLINIC MANAGEMENT END 


        ]
    },


    // ===========================================

    {
        path: '/:pathMatch(.*)*',
        redirect: '/404'
    }
)