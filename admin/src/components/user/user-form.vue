<!-- components/user/user-form.vue -->
<script setup lang="ts">
import { ref } from 'vue';
import type { FormInst, FormRules } from 'naive-ui';

const formRef = ref<FormInst | null>(null);
const formValue = ref({
  first_name: '',
  last_name: '',
  email: '',
  mobile: '',
  user_type: 'USER'
});

const rules: FormRules = {
  first_name: {
    required: true,
    message: $t('user.firstNameRequired')
  },
  email: {
    required: true,
    type: 'email',
    message: $t('user.emailInvalid')
  }
  // Add other validation rules
};

defineExpose({
  formRef,
  formValue
});
</script>

<template>
  <NForm
    ref="formRef"
    :model="formValue"
    :rules="rules"
    label-placement="left"
    label-width="100"
    require-mark-placement="right-hanging"
  >
    <NFormItem :label="$t('user.firstName')" path="first_name">
      <NInput v-model:value="formValue.first_name" />
    </NFormItem>
    <NFormItem :label="$t('user.lastName')" path="last_name">
      <NInput v-model:value="formValue.last_name" />
    </NFormItem>
    <NFormItem :label="$t('user.email')" path="email">
      <NInput v-model:value="formValue.email" />
    </NFormItem>
    <NFormItem :label="$t('user.mobile')" path="mobile">
      <NInput v-model:value="formValue.mobile" />
    </NFormItem>
    <NFormItem :label="$t('user.userType')" path="user_type">
      <NSelect
        v-model:value="formValue.user_type"
        :options="[
          { label: 'Super Admin', value: 'SUPERADMIN' },
          { label: 'User', value: 'USER' }
        ]"
      />
    </NFormItem>
  </NForm>
</template>
