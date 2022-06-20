<template>
    <BasicContent>
        <div class="row">
            <div class="col-3">
                <UserInfo
                    @followEvent="follow"
                    @unfollowEvent="unfollow"
                    :user="user"
                />
                <UserPostWrite @pushPost="pushPost" />
            </div>
            <div class="col-9">
                <UserPosts :posts="posts" />
            </div>
        </div>
    </BasicContent>
</template>

<script>
import BasicContent from "@/components/BasicContent";
import UserInfo from "@/components/UserInfo";
import UserPosts from "@/components/UserPosts";
import { reactive } from "vue";
import UserPostWrite from "@/components/UserPostWrite";
import { useRoute } from "vue-router"; // 使用自定义参数

export default {
    name: "UserSpace",
    components: {
        BasicContent,
        UserInfo,
        UserPosts,
        UserPostWrite,
    },

    setup() {
        const route = useRoute();
        console.log(route.params.userId); // 取出 url 中得对应参数的 value

        const user = reactive({
            userid: 1,
            username: "admin",
            // firstName: "ad",
            // lastName: "min",
            followerCount: 0,
            isFollowed: false,
        });

        const posts = reactive({
            count: 3,
            posts: [
                {
                    id: 1,
                    userId: 1,
                    content: "Hello fw",
                },
                {
                    id: 2,
                    userId: 1,
                    content: "Hello yxc",
                },
                {
                    id: 3,
                    userId: 1,
                    content: "Hello fd",
                },
            ],
        });

        const follow = () => {
            if (user.isFollowed) {
                return;
            }
            user.isFollowed = true;
            user.followerCount++;
        };

        const unfollow = () => {
            if (user.isFollowed === false) {
                return;
            }
            user.isFollowed = false;
            user.followerCount--;
        };

        const pushPost = (content) => {
            posts.count++;
            posts.posts.unshift({
                id: posts.count,
                userId: 1,
                content: content,
            });
        };

        return {
            user, // key value 相同可以简写为 user,
            posts,
            follow,
            unfollow,
            pushPost,
        };
    },
};
</script>

<style scoped></style>
