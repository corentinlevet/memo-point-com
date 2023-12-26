<template>
  <div class="container mt-4">
    <h2 class="mb-4">Memos</h2>

    <button class="btn btn-primary mb-3" @click="toggleForm">
      {{ showForm ? 'Masquer le formulaire' : 'Ajouter un memo' }}
    </button>

    <div v-if="showForm">
      <form>
        <div class="mb-3">
          <label class="form-label">Titre :</label>
          <input type="text" class="form-control" v-model="newMemoTitle" />
        </div>
        <div class="mb-3">
          <label class="form-label">Contenu :</label>
          <textarea class="form-control" v-model="newMemoContent" style="height: 150px"></textarea>
        </div>
        <button type="button" class="btn" :class="{ 'btn-info': selectedMemo !== null, 'btn-success': selectedMemo === null }" @click="addMemo">
          {{ selectedMemo !== null ? 'Mettre à jour la memo' : 'Ajouter un memo' }}
        </button>
      </form>
    </div>

    <div v-else>
      <ul class="list-group">
        <li class="list-group-item" v-for="(memo, index) in memos" :key="index">
          <h4>{{ memo.title }}</h4>
          <p style="word-break: break-all">{{ memo.content }}</p>
          <button class="btn btn-info me-2" @click="selectMemoForEdit(index)">Modifier</button>
          <button class="btn btn-danger" @click="confirmAndDelete(index)">Supprimer</button>
        </li>
      </ul>
    </div>
  </div>
</template>

<script lang="ts">
  import goServer from '@/api/go-server';

  export default {
    data() {
      return {
        memos: [],
        newMemoTitle: '',
        newMemoContent: '',
        showForm: false,
        selectedMemo: null,
      };
    },
    methods: {
      toggleForm() {
        this.showForm = !this.showForm;
        this.selectedMemo = null;
        this.newMemoTitle = '';
        this.newMemoContent = '';
      },
      addMemo() {
        if (this.newMemoTitle.trim() !== '' && this.newMemoContent.trim() !== '') {
          const newMemo = {
            id: null,
            title: this.newMemoTitle,
            content: this.newMemoContent,
          };

          if (this.selectedMemo !== null) {
            goServer.updateMemo(this.memos[this.selectedMemo].id, this.newMemoTitle, this.newMemoContent).then((response) => {
              newMemo.id = response.data.memo_id;
              this.memos[this.selectedMemo] = newMemo;
              this.selectedMemo = null;
              console.log(response.data.message);
            }).catch((error) => {
              console.log(error);
            });
          } else {
            goServer.createMemo(this.newMemoTitle, this.newMemoContent).then((response) => {
              newMemo.id = response.data.memo_id;
              this.memos.push(newMemo);
              console.log(response.data.message);
            }).catch((error) => {
              console.log(error);
            });
          }

          this.newMemoTitle = '';
          this.newMemoContent = '';
          this.showForm = false;
        }
      },
      selectMemoForEdit(index: number) {
        this.selectedMemo = index;
        this.newMemoTitle = this.memos[index].title;
        this.newMemoContent = this.memos[index].content;
        this.showForm = true;
      },
      confirmAndDelete(index: number) {
        const isConfirmed = window.confirm('Êtes-vous sûr de vouloir supprimer cette memo?');

        if (isConfirmed) {
          this.deleteMemo(index);
        }
      },
      deleteMemo(index: number) {
        const memoID = this.memos[index].id;

        goServer.deleteMemo(memoID).then((response) => {
          this.memos.splice(index, 1);
          console.log(response.data.message);
        }).catch((error) => {
          console.log(error);
        });
      },
      fetchMemos() {
        goServer.getMemos().then((response) => {
          const memos = response.data.memos;
          this.memos = memos.map((memo: any) => {
            return {
              id: memo.id,
              title: memo.title,
              content: memo.content,
            };
          });
        }).catch((error) => {
          console.log(error);
        });
      },
    },
    mounted() {
      this.fetchMemos();
    },
  };
</script>
