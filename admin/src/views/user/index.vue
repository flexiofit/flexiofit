<!-- <!-- views/user/index.vue --> -->
<script setup lang="ts">
import { ref, onMounted, h } from 'vue';
import { useUserStore } from '@/store/modules/users';
import { storeToRefs } from 'pinia';
import { NDataTable, NCard, NModal, NButton, NSpace } from 'naive-ui';
import type { DataTableColumns } from 'naive-ui';
import { $t } from '@/locales';
import type { User } from '@/store/modules/users/types';

const userStore = useUserStore();
const { users, loading } = storeToRefs(userStore);

interface TableColumn {
  title: string;
  key: string;
  checked: boolean;
  render?: (row: User) => any;
}

// Define table columns
const columns = ref<TableColumn[]>([
  {
    title: 'ID',
    key: 'id',
    checked: true
  },
  {
    title: $t('user.firstName'),
    key: 'first_name',
    checked: true
  },
  {
    title: $t('user.lastName'),
    key: 'last_name',
    checked: true
  },
  {
    title: $t('user.email'),
    key: 'email',
    checked: true
  },
  {
    title: $t('user.mobile'),
    key: 'mobile',
    checked: true
  },
  {
    title: $t('user.userType'),
    key: 'user_type',
    checked: true
  },
  {
    title: $t('common.action'),
    key: 'actions',
    checked: true,
    render: (row: User) => {
      return h(NSpace, {}, {
        default: () => [
          h(NButton, {
            size: 'small',
            type: 'primary',
            ghost: true,
            onClick: () => handleEdit(row.id)
          }, { default: () => $t('common.edit') }),
          h(NButton, {
            size: 'small',
            type: 'error',
            ghost: true,
            onClick: () => handleDelete(row.id)
          }, { default: () => $t('common.delete') })
        ]
      });
    }
  }
]);

// Selected rows for batch operations
const checkedRowKeys = ref<(string | number)[]>([]);

// Add user modal
const showAddModal = ref(false);

const handleAdd = () => {
  showAddModal.value = true;
};

const handleEdit = async (id: number) => {
  try {
    await userStore.fetchUserById(id.toString());
    // Implement your edit logic
  } catch (err) {
    if (err instanceof Error) {
      window.$message.error(err.message);
    } else {
      window.$message.error('Unknown error occurred');
    }
  }
};

const handleDelete = async (id: number) => {
  try {
    await userStore.removeUser(id.toString());
    window.$message.success($t('common.deleteSuccess'));
  } catch (err) {
    if (err instanceof Error) {
      window.$message.error(err.message);
    } else {
      window.$message.error('Unknown error occurred');
    }
  }
};

const handleBatchDelete = async () => {
  try {
    await Promise.all(
      checkedRowKeys.value.map(id => userStore.removeUser(id.toString()))
    );
    checkedRowKeys.value = [];
    window.$message.success($t('common.batchDeleteSuccess'));
  } catch (err) {
    if (err instanceof Error) {
      window.$message.error(err.message);
    } else {
      window.$message.error('Unknown error occurred');
    }
  }
};

const refresh = async () => {
  try {
    await userStore.fetchUsers();
    window.$message.success($t('common.refreshSuccess'));
  } catch (err) {
    if (err instanceof Error) {
      window.$message.error(err.message);
    } else {
      window.$message.error('Unknown error occurred');
    }
  }
};

onMounted(() => {
  refresh();
});
</script>

<template>
  <div>
    <NCard>
      <TableHeaderOperation
        v-model:columns="columns"
        :disabled-delete="!checkedRowKeys.length"
        :loading="loading"
        @add="handleAdd"
        @delete="handleBatchDelete"
        @refresh="refresh"
      >
        <!-- You can add custom buttons in the slots if needed -->
        <template #prefix>
          <!-- Additional prefix buttons -->
        </template>
        <template #suffix>
          <!-- Additional suffix buttons -->
        </template>
      </TableHeaderOperation>
      <NDataTable
        :loading="loading"
        :columns="columns.filter(col => col.checked)"
        :data="users"
        :row-key="(row: User) => row.id"
        v-model:checked-row-keys="checkedRowKeys"
        :scroll-x="1200"
      />
    </NCard>
    <!-- Add/Edit User Modal -->
    <NModal v-model:show="showAddModal" preset="card" :title="$t('user.addUser')">
      <!-- Add your form component here -->
    </NModal>
  </div>
</template>

<style scoped>
.n-card {
  margin: 16px;
}
</style>
