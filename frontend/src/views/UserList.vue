<template>
    <BasicContent>
        <div v-for="user in users" :key="user.id" class="card">
            <div class="card-body">
                <div class="row">
                    <div class="col-1">
                        <img class="img-fluid" :src="user.photo" alt="" />
                    </div>
                    <div class="col-11">
                        <div class="username">
                            {{ user.username }}
                        </div>
                        <div class="follower-count">
                            {{ user.followerCount }}
                        </div>
                    </div>
                </div>
            </div>
        </div></BasicContent
    >
</template>

<script>
import BasicContent from "@/components/BasicContent";
import $ from "jquery";
import { ref } from "vue";

export default {
    name: "UserList",
    components: {
        BasicContent,
    },

    setup() {
        let users = ref([]);

        $.ajax({
            url: "https://app165.acapp.acwing.com.cn/myspace/userlist/",
            type: "get",
            success(resp) {
                users.value = resp;
            },
        });

        return {
            users,
        };
    },
};
</script>

<style scoped>
img {
    border-radius: 50%;
}

.username {
    font-weight: bold;
    height: 50%;
}

.follower-count {
    font-size: 12px;
    color: gray;
    height: 50%;
}

.card {
    margin-bottom: 20px;
    cursor: pointer;
}

.card:hover {
    box-shadow: 2px 2px 10px lightgrey;
    transition: 500ms;
}

.img-field {
    display: flex;
    flex-direction: column;
    justify-content: center;
}
</style>
