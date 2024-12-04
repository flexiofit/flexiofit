<!-- <template> -->
<!--   <n-space vertical :size="12"> -->
<!--     <n-space> -->
<!--       <n-button @click="sortName">Sort By Name (Ascend)</n-button> -->
<!--       <n-button @click="filterAddress">Filter Address (London)</n-button> -->
<!--       <n-button @click="clearFilters">Clear Filters</n-button> -->
<!--       <n-button @click="clearSorter">Clear Sorter</n-button> -->
<!--     </n-space> -->
<!--     <n-data-table -->
<!--       ref="dataTableInst" -->
<!--       :columns="columns" -->
<!--       :data="data" -->
<!--       :pagination="pagination" -->
<!--     /> -->
<!--   </n-space> -->
<!-- </template> -->
<!---->
<!-- <script> -->
<!-- import { defineComponent, ref } from 'vue' -->
<!---->
<!-- const columns = [ -->
<!--   { -->
<!--     title: 'Name', -->
<!--     key: 'name' -->
<!--   }, -->
<!--   { -->
<!--     title: 'Age', -->
<!--     key: 'age', -->
<!--     sorter: (row1, row2) => row1.age - row2.age -->
<!--   }, -->
<!--   { -->
<!--     title: 'Chinese Score', -->
<!--     key: 'chinese', -->
<!--     defaultSortOrder: false, -->
<!--     sorter: { -->
<!--       compare: (a, b) => a.chinese - b.chinese, -->
<!--       multiple: 3 -->
<!--     } -->
<!--   }, -->
<!--   { -->
<!--     title: 'Math Score', -->
<!--     defaultSortOrder: false, -->
<!--     key: 'math', -->
<!--     sorter: { -->
<!--       compare: (a, b) => a.math - b.math, -->
<!--       multiple: 2 -->
<!--     } -->
<!--   }, -->
<!--   { -->
<!--     title: 'English Score', -->
<!--     defaultSortOrder: false, -->
<!--     key: 'english', -->
<!--     sorter: { -->
<!--       compare: (a, b) => a.english - b.english, -->
<!--       multiple: 1 -->
<!--     } -->
<!--   }, -->
<!--   { -->
<!--     title: 'Address', -->
<!--     key: 'address', -->
<!--     filterOptions: [ -->
<!--       { -->
<!--         label: 'London', -->
<!--         value: 'London' -->
<!--       }, -->
<!--       { -->
<!--         label: 'New York', -->
<!--         value: 'New York' -->
<!--       } -->
<!--     ], -->
<!--     filter(value, row) { -->
<!--       return ~row.address.indexOf(value) -->
<!--     } -->
<!--   } -->
<!-- ] -->
<!---->
<!-- const data = [ -->
<!--   { -->
<!--     key: 0, -->
<!--     name: 'John Brown', -->
<!--     age: 32, -->
<!--     address: 'New York No. 1 Lake Park', -->
<!--     chinese: 98, -->
<!--     math: 60, -->
<!--     english: 70 -->
<!--   }, -->
<!--   { -->
<!--     key: 1, -->
<!--     name: 'Jim Green', -->
<!--     age: 42, -->
<!--     address: 'London No. 1 Lake Park', -->
<!--     chinese: 98, -->
<!--     math: 66, -->
<!--     english: 89 -->
<!--   }, -->
<!--   { -->
<!--     key: 2, -->
<!--     name: 'Joe Black', -->
<!--     age: 32, -->
<!--     address: 'Sidney No. 1 Lake Park', -->
<!--     chinese: 98, -->
<!--     math: 66, -->
<!--     english: 89 -->
<!--   }, -->
<!--   { -->
<!--     key: 3, -->
<!--     name: 'Jim Red', -->
<!--     age: 32, -->
<!--     address: 'London No. 2 Lake Park', -->
<!--     chinese: 88, -->
<!--     math: 99, -->
<!--     english: 89 -->
<!--   } -->
<!-- ] -->
<!---->
<!-- export default defineComponent({ -->
<!--   setup() { -->
<!--     const dataTableInstRef = ref(null) -->
<!--     return { -->
<!--       data, -->
<!--       columns, -->
<!--       dataTableInst: dataTableInstRef, -->
<!--       pagination: ref({ pageSize: 5 }), -->
<!--       filterAddress() { -->
<!--         dataTableInstRef.value.filter({ -->
<!--           address: ['London'] -->
<!--         }) -->
<!--       }, -->
<!--       sortName() { -->
<!--         dataTableInstRef.value.sort('name', 'ascend') -->
<!--       }, -->
<!--       clearFilters() { -->
<!--         dataTableInstRef.value.filter(null) -->
<!--       }, -->
<!--       clearSorter() { -->
<!--         dataTableInstRef.value.sort(null) -->
<!--       } -->
<!--     } -->
<!--   } -->
<!-- }) -->
<!-- </script> -->

<template>
  <n-space vertical :size="12">
    <n-space>
      <n-button @click="sortName">Sort By Name (Ascend)</n-button>
      <n-button @click="filterAddress">Filter Address (London)</n-button>
      <n-button @click="clearFilters">Clear Filters</n-button>
      <n-button @click="clearSorter">Clear Sorter</n-button>
    </n-space>

    <!-- Data Table -->
    <n-data-table
      ref="dataTableInst"
      :columns="columns"
      :data="data"
      :pagination="pagination"
      :loading="loading"
    />
  </n-space>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue';
import { fetchAllUsers, fetchUserById, createUser, updateUser, deleteUser } from '@/service/api/user';
import { NButton, NDataTable, NSpace } from 'naive-ui';

export default defineComponent({
  name: 'UserManagement',
  setup() {
    const dataTableInstRef = ref(null);
    const pagination = ref({ pageSize: 5, current: 1 });
    const data = ref([]);
    const loading = ref(false);
    const columns = [
      {
        title: 'Name',
        key: 'name',
        sorter: (row1, row2) => row1.name.localeCompare(row2.name)
      },
      {
        title: 'Age',
        key: 'age',
        sorter: (row1, row2) => row1.age - row2.age
      },
      {
        title: 'Address',
        key: 'address',
        filterOptions: [
          { label: 'London', value: 'London' },
          { label: 'New York', value: 'New York' },
        ],
        filter(value, row) {
          return row.address.includes(value);
        }
      }
    ];

    // Fetch all users initially on component mount
    onMounted(async () => {
      await loadUsers();
    });

    // Fetch users from API
    const loadUsers = async () => {
      loading.value = true;
      try {
        const response = await fetchAllUsers();
        data.value = response;
      } catch (error) {
        console.error('Failed to load users', error);
      } finally {
        loading.value = false;
      }
    };

    // Sort users by Name
    const sortName = () => {
      dataTableInstRef.value.sort('name', 'ascend');
    };

    // Filter users by address
    const filterAddress = () => {
      dataTableInstRef.value.filter({
        address: ['London']
      });
    };

    // Clear all filters
    const clearFilters = () => {
      dataTableInstRef.value.filter(null);
    };

    // Clear sorting
    const clearSorter = () => {
      dataTableInstRef.value.sort(null);
    };

    // Add new user (Example action)
    const addUser = async (user) => {
      loading.value = true;
      try {
        const createdUser = await createUser(user);
        data.value.push(createdUser);
      } catch (error) {
        console.error('Failed to create user', error);
      } finally {
        loading.value = false;
      }
    };

    // Update an existing user (Example action)
    const updateUserDetails = async (id, user) => {
      loading.value = true;
      try {
        const updatedUser = await updateUser(id, user);
        const index = data.value.findIndex((item) => item.id === id);
        if (index !== -1) {
          data.value[index] = updatedUser;
        }
      } catch (error) {
        console.error('Failed to update user', error);
      } finally {
        loading.value = false;
      }
    };

    // Delete a user (Example action)
    const deleteUserDetails = async (id) => {
      loading.value = true;
      try {
        await deleteUser(id);
        data.value = data.value.filter((user) => user.id !== id);
      } catch (error) {
        console.error('Failed to delete user', error);
      } finally {
        loading.value = false;
      }
    };

    return {
      data,
      columns,
      pagination,
      loading,
      dataTableInst: dataTableInstRef,
      loadUsers,
      sortName,
      filterAddress,
      clearFilters,
      clearSorter,
      addUser,
      updateUserDetails,
      deleteUserDetails
    };
  }
});
</script>


