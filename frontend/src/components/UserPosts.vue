<template>
    <div class="card">
        <div class="card-body">
            <div v-for="post in posts.posts" :key="post.id">
                <div class="card post-item">
                    <div class="card-body">
                        {{ post.content }}
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "UserPosts",
    props: {
        posts: {
            type: Object,
            required: true,
        },
    },
    components: {},
};
</script>

<style scoped>
.post-item {
    margin-bottom: 10px;
}
</style>
