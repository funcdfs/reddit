<template>
    <div class="card edit-field">
        <div class="card-body">
            <div class="mb-3">
                <label for="exampleFormControlTextarea1" class="form-label"
                    >new post below:</label
                >
                <textarea
                    v-model="content"
                    class="form-control"
                    id="edit-post"
                    rows="3"
                ></textarea>
            </div>
        </div>
        <div>
            {{ content }}
        </div>
        <button @click="pushPost" type="button" class="btn btn-primary">
            push
        </button>
    </div>
</template>

<script>
import { ref } from "vue";

export default {
    name: "UserPostWrite",
    components: {},

    setup(_, context) {
        let content = ref("");

        const pushPost = () => {
            context.emit("pushPost", content.value);
            content.value = "";
        };

        return {
            content,
            pushPost,
        };
    },
};
</script>

<style scoped>
.edit-field {
    margin-top: 10px;
}
</style>
