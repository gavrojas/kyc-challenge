import { createRouter, createWebHistory } from 'vue-router'
import OnboardingStep1 from '@/views/OnboardingStep1.vue';
import OnboardingStep2 from '@/views/OnboardingStep2.vue';
import OnboardingStep3 from '@/views/OnboardingStep3.vue';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/', redirect: '/onboarding-step-1',
    },
    { path: '/onboarding-step-1', component: OnboardingStep1 },
    { path: '/onboarding-step-2', component: OnboardingStep2 },
    { path: '/onboarding-step-3', component: OnboardingStep3 },
  ]
})

export default router
