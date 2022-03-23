package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"unsafe"
)

var data = `
{
    "err_no": 0,
    "err_msg": "success",
    "data": [
        {
            "article_id": "7051406419823689765",
            "article_info": {
                "article_id": "7051406419823689765",
                "user_id": "2348212566892574",
                "category_id": "6809637769959178254",
                "tag_ids": [
                    6809640364677267000,
                    6809640620030689000
                ],
                "visible_level": 0,
                "link_url": "",
                "cover_image": "",
                "is_gfw": 0,
                "title": "Go 中实现用户的每日限额（比如一天只能领三次福利）",
                "brief_content": "本文主要讲解固定时间窗口限流算法，使用场景比如： 每个手机号每天只能发5条验证码短信 每个用户每小时只能连续尝试3次密码 每个会员每天只能领3次福利",
                "is_english": 0,
                "is_original": 1,
                "user_index": 7.372922110335837,
                "original_type": 0,
                "original_author": "",
                "content": "",
                "ctime": "1641783617",
                "mtime": "1641785133",
                "rtime": "1641785133",
                "draft_id": "7051406156576587806",
                "view_count": 79,
                "collect_count": 1,
                "digg_count": 0,
                "comment_count": 0,
                "hot_index": 3,
                "is_hot": 0,
                "rank_index": 0.9675429,
                "status": 2,
                "verify_status": 1,
                "audit_status": 2,
                "mark_content": "",
                "display_count": 0
            },
            "author_user_info": {
                "user_id": "2348212566892574",
                "user_name": "万俊峰Kevin",
                "company": "七牛云",
                "job_title": "技术副总裁",
                "avatar_large": "https://p9-passport.byteacctimg.com/img/user-avatar/5f7368678ecd20fc46d6e85f860b4a24~300x300.image",
                "level": 3,
                "description": "go-zero作者\n七牛云技术副总裁\n阿里云MVP\nArchSummit明星讲师\nGopherChina金牌讲师\nQCon+出品人兼讲师\n腾讯云开发者大会讲师",
                "followee_count": 11,
                "follower_count": 434,
                "post_article_count": 73,
                "digg_article_count": 76,
                "got_digg_count": 613,
                "got_view_count": 80404,
                "post_shortmsg_count": 3,
                "digg_shortmsg_count": 1,
                "isfollowed": true,
                "favorable_author": 0,
                "power": 1433,
                "study_point": 0,
                "university": {
                    "university_id": "0",
                    "name": "",
                    "logo": ""
                },
                "major": {
                    "major_id": "0",
                    "parent_id": "0",
                    "name": ""
                },
                "student_status": 0,
                "select_event_count": 0,
                "select_online_course_count": 0,
                "identity": 0,
                "is_select_annual": true,
                "select_annual_rank": 0,
                "annual_list_type": 0,
                "extraMap": {},
                "is_logout": 0
            },
            "category": {
                "category_id": "6809637769959178254",
                "category_name": "后端",
                "category_url": "backend",
                "rank": 1,
                "back_ground": "https://lc-mhke0kuv.cn-n1.lcfile.com/fb3b208d06e6fe32.png",
                "icon": "https://lc-mhke0kuv.cn-n1.lcfile.com/a2ec01b816abd4c5.png",
                "ctime": 1457483880,
                "mtime": 1432503193,
                "show_type": 3,
                "item_type": 2,
                "promote_tag_cap": 4,
                "promote_priority": 1
            },
            "tags": [
                {
                    "id": 2546494,
                    "tag_id": "6809640364677267469",
                    "tag_name": "Go",
                    "color": "#64D7E3",
                    "icon": "https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/leancloud-assets/1aae38ab22d12b654cfa.png~tplv-t2oaga2asx-image.image",
                    "back_ground": "",
                    "show_navi": 0,
                    "ctime": 1432234497,
                    "mtime": 1641802500,
                    "id_type": 9,
                    "tag_alias": "",
                    "post_article_count": 10835,
                    "concern_user_count": 88731
                },
                {
                    "id": 2546679,
                    "tag_id": "6809640620030689287",
                    "tag_name": "微服务",
                    "color": "#000000",
                    "icon": "https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/leancloud-assets/e5d6c745cae839047dc7.png~tplv-t2oaga2asx-image.image",
                    "back_ground": "",
                    "show_navi": 0,
                    "ctime": 1457382673,
                    "mtime": 1641803418,
                    "id_type": 9,
                    "tag_alias": "",
                    "post_article_count": 3552,
                    "concern_user_count": 60435
                }
            ],
            "user_interact": {
                "id": 7051406419823690000,
                "omitempty": 2,
                "user_id": 3984285869016509,
                "is_digg": false,
                "is_follow": false,
                "is_collect": false
            },
            "org": {
                "org_info": null,
                "org_user": null,
                "is_followed": false
            },
            "req_id": "202201101640270102080650252144A87B",
            "status": {
                "push_status": 0
            }
        },
        {
            "article_id": "7051370329955893278",
            "article_info": {
                "article_id": "7051370329955893278",
                "user_id": "1574156384091320",
                "category_id": "6809637771511070734",
                "tag_ids": [
                    6809640375880253000,
                    6809640419505209000
                ],
                "visible_level": 0,
                "link_url": "",
                "cover_image": "https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/f76a01216e9e4d8d9c53ca907a14130c~tplv-k3u1fbpfcp-watermark.image?",
                "is_gfw": 0,
                "title": "GitHub 公布 2021 Top 10 博文「GitHub 热点速览」",
                "brief_content": "2021 年在这周彻底同我们告别了，在本周的「News 快读」模块你可以看到过去一年 GitHub 的热门文章，其中有我们熟悉的可能让很多程序员“失业”的 Copilot，还有官方的云端 IDE Co",
                "is_english": 0,
                "is_original": 1,
                "user_index": 9.31103982378772,
                "original_type": 0,
                "original_author": "",
                "content": "",
                "ctime": "1641775126",
                "mtime": "1641775633",
                "rtime": "1641775632",
                "draft_id": "7051163788560564231",
                "view_count": 326,
                "collect_count": 0,
                "digg_count": 0,
                "comment_count": 0,
                "hot_index": 16,
                "is_hot": 0,
                "rank_index": 1.48891761,
                "status": 2,
                "verify_status": 1,
                "audit_status": 2,
                "mark_content": "",
                "display_count": 0
            },
            "author_user_info": {
                "user_id": "1574156384091320",
                "user_name": "HelloGitHub",
                "company": "公众号：HelloGitHub",
                "job_title": "",
                "avatar_large": "https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/gold-user-assets/2019/3/28/169c31b06dcd6e90~tplv-t2oaga2asx-image.image",
                "level": 5,
                "description": "找开源项目可以试试 HelloGitHub 小程序～",
                "followee_count": 33,
                "follower_count": 10802,
                "post_article_count": 376,
                "digg_article_count": 264,
                "got_digg_count": 9118,
                "got_view_count": 810970,
                "post_shortmsg_count": 1087,
                "digg_shortmsg_count": 220,
                "isfollowed": true,
                "favorable_author": 1,
                "power": 16363,
                "study_point": 0,
                "university": {
                    "university_id": "0",
                    "name": "",
                    "logo": ""
                },
                "major": {
                    "major_id": "0",
                    "parent_id": "0",
                    "name": ""
                },
                "student_status": 0,
                "select_event_count": 0,
                "select_online_course_count": 0,
                "identity": 0,
                "is_select_annual": true,
                "select_annual_rank": 0,
                "annual_list_type": 0,
                "extraMap": {},
                "is_logout": 0
            },
            "category": {
                "category_id": "6809637771511070734",
                "category_name": "开发工具",
                "category_url": "freebie",
                "rank": 6,
                "back_ground": "https://lc-mhke0kuv.cn-n1.lcfile.com/412957a61f414c0b.png",
                "icon": "https://lc-mhke0kuv.cn-n1.lcfile.com/eb20ab1334d9abea.png",
                "ctime": 1457483920,
                "mtime": 1432503202,
                "show_type": 3,
                "item_type": 2,
                "promote_tag_cap": 4,
                "promote_priority": 6
            },
            "tags": [
                {
                    "id": 2546502,
                    "tag_id": "6809640375880253447",
                    "tag_name": "GitHub",
                    "color": "#616161",
                    "icon": "https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/leancloud-assets/0d614af263aa63aa6a77.png~tplv-t2oaga2asx-image.image",
                    "back_ground": "",
                    "show_navi": 0,
                    "ctime": 1432234558,
                    "mtime": 1641803412,
                    "id_type": 9,
                    "tag_alias": "",
                    "post_article_count": 11281,
                    "concern_user_count": 397818
                },
                {
                    "id": 2546535,
                    "tag_id": "6809640419505209358",
                    "tag_name": "开源",
                    "color": "#6EBD68",
                    "icon": "https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/leancloud-assets/553ecacd498946a9a6d9.png~tplv-t2oaga2asx-image.image",
                    "back_ground": "",
                    "show_navi": 0,
                    "ctime": 1435972427,
                    "mtime": 1641803418,
                    "id_type": 9,
                    "tag_alias": "",
                    "post_article_count": 6377,
                    "concern_user_count": 220213
                }
            ],
            "user_interact": {
                "id": 7051370329955893000,
                "omitempty": 2,
                "user_id": 3984285869016509,
                "is_digg": false,
                "is_follow": false,
                "is_collect": false
            },
            "org": {
                "org_info": null,
                "org_user": null,
                "is_followed": false
            },
            "req_id": "202201101640270102080650252144A87B",
            "status": {
                "push_status": 0
            }
        },
        {
            "article_id": "7051406419823689765",
            "article_info": {
                "article_id": "7051406419823689765",
                "user_id": "2348212566892574",
                "category_id": "6809637769959178254",
                "tag_ids": [
                    6809640364677267000,
                    6809640620030689000
                ],
                "visible_level": 0,
                "link_url": "",
                "cover_image": "",
                "is_gfw": 0,
                "title": "Go 中实现用户的每日限额（比如一天只能领三次福利）",
                "brief_content": "本文主要讲解固定时间窗口限流算法，使用场景比如： 每个手机号每天只能发5条验证码短信 每个用户每小时只能连续尝试3次密码 每个会员每天只能领3次福利",
                "is_english": 0,
                "is_original": 1,
                "user_index": 7.372922110335837,
                "original_type": 0,
                "original_author": "",
                "content": "",
                "ctime": "1641783617",
                "mtime": "1641785133",
                "rtime": "1641785133",
                "draft_id": "7051406156576587806",
                "view_count": 79,
                "collect_count": 1,
                "digg_count": 0,
                "comment_count": 0,
                "hot_index": 3,
                "is_hot": 0,
                "rank_index": 0.9675429,
                "status": 2,
                "verify_status": 1,
                "audit_status": 2,
                "mark_content": "",
                "display_count": 0
            },
            "author_user_info": {
                "user_id": "2348212566892574",
                "user_name": "万俊峰Kevin",
                "company": "七牛云",
                "job_title": "技术副总裁",
                "avatar_large": "https://p9-passport.byteacctimg.com/img/user-avatar/5f7368678ecd20fc46d6e85f860b4a24~300x300.image",
                "level": 3,
                "description": "go-zero作者\n七牛云技术副总裁\n阿里云MVP\nArchSummit明星讲师\nGopherChina金牌讲师\nQCon+出品人兼讲师\n腾讯云开发者大会讲师",
                "followee_count": 11,
                "follower_count": 434,
                "post_article_count": 73,
                "digg_article_count": 76,
                "got_digg_count": 613,
                "got_view_count": 80404,
                "post_shortmsg_count": 3,
                "digg_shortmsg_count": 1,
                "isfollowed": true,
                "favorable_author": 0,
                "power": 1433,
                "study_point": 0,
                "university": {
                    "university_id": "0",
                    "name": "",
                    "logo": ""
                },
                "major": {
                    "major_id": "0",
                    "parent_id": "0",
                    "name": ""
                },
                "student_status": 0,
                "select_event_count": 0,
                "select_online_course_count": 0,
                "identity": 0,
                "is_select_annual": true,
                "select_annual_rank": 0,
                "annual_list_type": 0,
                "extraMap": {},
                "is_logout": 0
            },
            "category": {
                "category_id": "6809637769959178254",
                "category_name": "后端",
                "category_url": "backend",
                "rank": 1,
                "back_ground": "https://lc-mhke0kuv.cn-n1.lcfile.com/fb3b208d06e6fe32.png",
                "icon": "https://lc-mhke0kuv.cn-n1.lcfile.com/a2ec01b816abd4c5.png",
                "ctime": 1457483880,
                "mtime": 1432503193,
                "show_type": 3,
                "item_type": 2,
                "promote_tag_cap": 4,
                "promote_priority": 1
            },
            "tags": [
                {
                    "id": 2546494,
                    "tag_id": "6809640364677267469",
                    "tag_name": "Go",
                    "color": "#64D7E3",
                    "icon": "https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/leancloud-assets/1aae38ab22d12b654cfa.png~tplv-t2oaga2asx-image.image",
                    "back_ground": "",
                    "show_navi": 0,
                    "ctime": 1432234497,
                    "mtime": 1641802500,
                    "id_type": 9,
                    "tag_alias": "",
                    "post_article_count": 10835,
                    "concern_user_count": 88731
                },
                {
                    "id": 2546679,
                    "tag_id": "6809640620030689287",
                    "tag_name": "微服务",
                    "color": "#000000",
                    "icon": "https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/leancloud-assets/e5d6c745cae839047dc7.png~tplv-t2oaga2asx-image.image",
                    "back_ground": "",
                    "show_navi": 0,
                    "ctime": 1457382673,
                    "mtime": 1641803418,
                    "id_type": 9,
                    "tag_alias": "",
                    "post_article_count": 3552,
                    "concern_user_count": 60435
                }
            ],
            "user_interact": {
                "id": 7051406419823690000,
                "omitempty": 2,
                "user_id": 3984285869016509,
                "is_digg": false,
                "is_follow": false,
                "is_collect": false
            },
            "org": {
                "org_info": null,
                "org_user": null,
                "is_followed": false
            },
            "req_id": "202201101640270102080650252144A87B",
            "status": {
                "push_status": 0
            }
        },
        {
            "article_id": "7051370329955893278",
            "article_info": {
                "article_id": "7051370329955893278",
                "user_id": "1574156384091320",
                "category_id": "6809637771511070734",
                "tag_ids": [
                    6809640375880253000,
                    6809640419505209000
                ],
                "visible_level": 0,
                "link_url": "",
                "cover_image": "https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/f76a01216e9e4d8d9c53ca907a14130c~tplv-k3u1fbpfcp-watermark.image?",
                "is_gfw": 0,
                "title": "GitHub 公布 2021 Top 10 博文「GitHub 热点速览」",
                "brief_content": "2021 年在这周彻底同我们告别了，在本周的「News 快读」模块你可以看到过去一年 GitHub 的热门文章，其中有我们熟悉的可能让很多程序员“失业”的 Copilot，还有官方的云端 IDE Co",
                "is_english": 0,
                "is_original": 1,
                "user_index": 9.31103982378772,
                "original_type": 0,
                "original_author": "",
                "content": "",
                "ctime": "1641775126",
                "mtime": "1641775633",
                "rtime": "1641775632",
                "draft_id": "7051163788560564231",
                "view_count": 326,
                "collect_count": 0,
                "digg_count": 0,
                "comment_count": 0,
                "hot_index": 16,
                "is_hot": 0,
                "rank_index": 1.48891761,
                "status": 2,
                "verify_status": 1,
                "audit_status": 2,
                "mark_content": "",
                "display_count": 0
            },
            "author_user_info": {
                "user_id": "1574156384091320",
                "user_name": "HelloGitHub",
                "company": "公众号：HelloGitHub",
                "job_title": "",
                "avatar_large": "https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/gold-user-assets/2019/3/28/169c31b06dcd6e90~tplv-t2oaga2asx-image.image",
                "level": 5,
                "description": "找开源项目可以试试 HelloGitHub 小程序～",
                "followee_count": 33,
                "follower_count": 10802,
                "post_article_count": 376,
                "digg_article_count": 264,
                "got_digg_count": 9118,
                "got_view_count": 810970,
                "post_shortmsg_count": 1087,
                "digg_shortmsg_count": 220,
                "isfollowed": true,
                "favorable_author": 1,
                "power": 16363,
                "study_point": 0,
                "university": {
                    "university_id": "0",
                    "name": "",
                    "logo": ""
                },
                "major": {
                    "major_id": "0",
                    "parent_id": "0",
                    "name": ""
                },
                "student_status": 0,
                "select_event_count": 0,
                "select_online_course_count": 0,
                "identity": 0,
                "is_select_annual": true,
                "select_annual_rank": 0,
                "annual_list_type": 0,
                "extraMap": {},
                "is_logout": 0
            },
            "category": {
                "category_id": "6809637771511070734",
                "category_name": "开发工具",
                "category_url": "freebie",
                "rank": 6,
                "back_ground": "https://lc-mhke0kuv.cn-n1.lcfile.com/412957a61f414c0b.png",
                "icon": "https://lc-mhke0kuv.cn-n1.lcfile.com/eb20ab1334d9abea.png",
                "ctime": 1457483920,
                "mtime": 1432503202,
                "show_type": 3,
                "item_type": 2,
                "promote_tag_cap": 4,
                "promote_priority": 6
            },
            "tags": [
                {
                    "id": 2546502,
                    "tag_id": "6809640375880253447",
                    "tag_name": "GitHub",
                    "color": "#616161",
                    "icon": "https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/leancloud-assets/0d614af263aa63aa6a77.png~tplv-t2oaga2asx-image.image",
                    "back_ground": "",
                    "show_navi": 0,
                    "ctime": 1432234558,
                    "mtime": 1641803412,
                    "id_type": 9,
                    "tag_alias": "",
                    "post_article_count": 11281,
                    "concern_user_count": 397818
                },
                {
                    "id": 2546535,
                    "tag_id": "6809640419505209358",
                    "tag_name": "开源",
                    "color": "#6EBD68",
                    "icon": "https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/leancloud-assets/553ecacd498946a9a6d9.png~tplv-t2oaga2asx-image.image",
                    "back_ground": "",
                    "show_navi": 0,
                    "ctime": 1435972427,
                    "mtime": 1641803418,
                    "id_type": 9,
                    "tag_alias": "",
                    "post_article_count": 6377,
                    "concern_user_count": 220213
                }
            ],
            "user_interact": {
                "id": 7051370329955893000,
                "omitempty": 2,
                "user_id": 3984285869016509,
                "is_digg": false,
                "is_follow": false,
                "is_collect": false
            },
            "org": {
                "org_info": null,
                "org_user": null,
                "is_followed": false
            },
            "req_id": "202201101640270102080650252144A87B",
            "status": {
                "push_status": 0
            }
        },
        {
            "article_id": "7050768976355262501",
            "article_info": {
                "article_id": "7050768976355262501",
                "user_id": "588993962972269",
                "category_id": "6809637769959178254",
                "tag_ids": [
                    6809640600502010000,
                    6809640371019055000
                ],
                "visible_level": 0,
                "link_url": "",
                "cover_image": "",
                "is_gfw": 0,
                "title": "数据库与缓存一致性方案分享",
                "brief_content": "做C端相关业务，目前主流的关系型数据库在高并发的查询请求场景下，很难做到低延迟的高并发，甚至有可能被打挂。因此引入缓存中间件是一个常见的解决方案，但如何保证缓存与数据库的一致性，便成为了一个棘手的问题",
                "is_english": 0,
                "is_original": 1,
                "user_index": 7.773422163647961,
                "original_type": 0,
                "original_author": "",
                "content": "",
                "ctime": "1641635107",
                "mtime": "1641637283",
                "rtime": "1641637283",
                "draft_id": "7050768448388874253",
                "view_count": 1439,
                "collect_count": 5,
                "digg_count": 4,
                "comment_count": 0,
                "hot_index": 75,
                "is_hot": 0,
                "rank_index": 0.55070508,
                "status": 2,
                "verify_status": 1,
                "audit_status": 2,
                "mark_content": "",
                "display_count": 0
            },
            "author_user_info": {
                "user_id": "588993962972269",
                "user_name": "呼呼虎",
                "company": "mt",
                "job_title": "后端工程师",
                "avatar_large": "https://p6-passport.byteacctimg.com/img/user-avatar/c6f6eb2675eff504f9fd46786baf36b6~300x300.image",
                "level": 2,
                "description": "java python 音视频处理",
                "followee_count": 19,
                "follower_count": 55,
                "post_article_count": 30,
                "digg_article_count": 45,
                "got_digg_count": 143,
                "got_view_count": 55522,
                "post_shortmsg_count": 0,
                "digg_shortmsg_count": 0,
                "isfollowed": true,
                "favorable_author": 0,
                "power": 698,
                "study_point": 0,
                "university": {
                    "university_id": "0",
                    "name": "",
                    "logo": ""
                },
                "major": {
                    "major_id": "0",
                    "parent_id": "0",
                    "name": ""
                },
                "student_status": 0,
                "select_event_count": 0,
                "select_online_course_count": 0,
                "identity": 0,
                "is_select_annual": false,
                "select_annual_rank": 0,
                "annual_list_type": 0,
                "extraMap": {},
                "is_logout": 0
            },
            "category": {
                "category_id": "6809637769959178254",
                "category_name": "后端",
                "category_url": "backend",
                "rank": 1,
                "back_ground": "https://lc-mhke0kuv.cn-n1.lcfile.com/fb3b208d06e6fe32.png",
                "icon": "https://lc-mhke0kuv.cn-n1.lcfile.com/a2ec01b816abd4c5.png",
                "ctime": 1457483880,
                "mtime": 1432503193,
                "show_type": 3,
                "item_type": 2,
                "promote_tag_cap": 4,
                "promote_priority": 1
            },
            "tags": [
                {
                    "id": 2546666,
                    "tag_id": "6809640600502009863",
                    "tag_name": "数据库",
                    "color": "#000000",
                    "icon": "https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/leancloud-assets/c57a5426c2a8ab565881.png~tplv-t2oaga2asx-image.image",
                    "back_ground": "",
                    "show_navi": 0,
                    "ctime": 1446683560,
                    "mtime": 1641803717,
                    "id_type": 9,
                    "tag_alias": "",
                    "post_article_count": 13193,
                    "concern_user_count": 289289
                },
                {
                    "id": 2546499,
                    "tag_id": "6809640371019055111",
                    "tag_name": "Redis",
                    "color": "#A51912",
                    "icon": "https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/leancloud-assets/4045af43b278afc7229b.png~tplv-t2oaga2asx-image.image",
                    "back_ground": "",
                    "show_navi": 0,
                    "ctime": 1432234546,
                    "mtime": 1641803990,
                    "id_type": 9,
                    "tag_alias": "",
                    "post_article_count": 7439,
                    "concern_user_count": 152789
                }
            ],
            "user_interact": {
                "id": 7050768976355262000,
                "omitempty": 2,
                "user_id": 3984285869016509,
                "is_digg": false,
                "is_follow": false,
                "is_collect": false
            },
            "org": {
                "org_info": null,
                "org_user": null,
                "is_followed": false
            },
            "req_id": "202201101640270102080650252144A87B",
            "status": {
                "push_status": 0
            }
        },
        {
            "article_id": "7050029639825096717",
            "article_info": {
                "article_id": "7050029639825096717",
                "user_id": "123560414944158",
                "category_id": "6809637769959178254",
                "tag_ids": [
                    6809640408797168000,
                    6809640495594078000
                ],
                "visible_level": 0,
                "link_url": "",
                "cover_image": "",
                "is_gfw": 0,
                "title": "Rust 12月学习资源",
                "brief_content": "聚焦学习 Rust 的网络资源 Rust 编译为什么慢？ 本文作者使用的硬件很高端，AMD Ryzen9 5950X CPU 32 核，128GB内存 和 SSD 硬盘。理论上编译 Rust 项目是非",
                "is_english": 0,
                "is_original": 1,
                "user_index": 5.896806271936629,
                "original_type": 0,
                "original_author": "",
                "content": "",
                "ctime": "1641462972",
                "mtime": "1641467457",
                "rtime": "-62135596800",
                "draft_id": "7050029200106848270",
                "view_count": 39,
                "collect_count": 0,
                "digg_count": 1,
                "comment_count": 0,
                "hot_index": 2,
                "is_hot": 0,
                "rank_index": 4e-8,
                "status": 1,
                "verify_status": 1,
                "audit_status": 2,
                "mark_content": "",
                "display_count": 0
            },
            "author_user_info": {
                "user_id": "123560414944158",
                "user_name": "Rust_Magazine",
                "company": "一本 Rust 语言开源杂志，每月末发刊",
                "job_title": "《Rust 中文精选》",
                "avatar_large": "https://p9-passport.byteacctimg.com/img/user-avatar/d0bd5f0426b8396a472f74eda0cf7808~300x300.image",
                "level": 3,
                "description": "定期更新 Rust 语言社区干货",
                "followee_count": 0,
                "follower_count": 463,
                "post_article_count": 198,
                "digg_article_count": 93,
                "got_digg_count": 557,
                "got_view_count": 162615,
                "post_shortmsg_count": 3,
                "digg_shortmsg_count": 1,
                "isfollowed": true,
                "favorable_author": 1,
                "power": 2185,
                "study_point": 0,
                "university": {
                    "university_id": "0",
                    "name": "",
                    "logo": ""
                },
                "major": {
                    "major_id": "0",
                    "parent_id": "0",
                    "name": ""
                },
                "student_status": 0,
                "select_event_count": 0,
                "select_online_course_count": 0,
                "identity": 0,
                "is_select_annual": false,
                "select_annual_rank": 0,
                "annual_list_type": 0,
                "extraMap": {},
                "is_logout": 0
            },
            "category": {
                "category_id": "6809637769959178254",
                "category_name": "后端",
                "category_url": "backend",
                "rank": 1,
                "back_ground": "https://lc-mhke0kuv.cn-n1.lcfile.com/fb3b208d06e6fe32.png",
                "icon": "https://lc-mhke0kuv.cn-n1.lcfile.com/a2ec01b816abd4c5.png",
                "ctime": 1457483880,
                "mtime": 1432503193,
                "show_type": 3,
                "item_type": 2,
                "promote_tag_cap": 4,
                "promote_priority": 1
            },
            "tags": [
                {
                    "id": 2546527,
                    "tag_id": "6809640408797167623",
                    "tag_name": "后端",
                    "color": "#C679FF",
                    "icon": "https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/leancloud-assets/d83da9d012ddb7ae85f4.png~tplv-t2oaga2asx-image.image",
                    "back_ground": "",
                    "show_navi": 1,
                    "ctime": 1435971556,
                    "mtime": 1641803746,
                    "id_type": 9,
                    "tag_alias": "",
                    "post_article_count": 95981,
                    "concern_user_count": 438297
                },
                {
                    "id": 2546590,
                    "tag_id": "6809640495594078216",
                    "tag_name": "Rust",
                    "color": "#616161",
                    "icon": "https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/leancloud-assets/01787a4f2859cde4db11.png~tplv-t2oaga2asx-image.image",
                    "back_ground": "",
                    "show_navi": 0,
                    "ctime": 1439498965,
                    "mtime": 1641801277,
                    "id_type": 9,
                    "tag_alias": "",
                    "post_article_count": 964,
                    "concern_user_count": 8662
                }
            ],
            "user_interact": {
                "id": 7050029639825097000,
                "omitempty": 2,
                "user_id": 3984285869016509,
                "is_digg": false,
                "is_follow": false,
                "is_collect": false
            },
            "org": {
                "org_info": null,
                "org_user": null,
                "is_followed": false
            },
            "req_id": "202201101640270102080650252144A87B",
            "status": {
                "push_status": 0
            }
        },
        {
            "article_id": "7050027158135111710",
            "article_info": {
                "article_id": "7050027158135111710",
                "user_id": "123560414944158",
                "category_id": "6809637769959178254",
                "tag_ids": [
                    6809640408797168000,
                    6809640495594078000
                ],
                "visible_level": 0,
                "link_url": "",
                "cover_image": "",
                "is_gfw": 0,
                "title": "Rust 社区 12月活动回顾",
                "brief_content": "后期编辑：张汉东 【线上】Rust 唠嗑室本月汇总 来源：Rust 唠嗑室 主持人：MikeTang 后期编辑：高宪凤 《Rust唠嗑室》第39期 - 来分享分享你们用Rust做的项目 时间: 202",
                "is_english": 0,
                "is_original": 1,
                "user_index": 5.909357461683179,
                "original_type": 0,
                "original_author": "",
                "content": "",
                "ctime": "1641462457",
                "mtime": "1641467430",
                "rtime": "-62135596800",
                "draft_id": "7050022974216208391",
                "view_count": 27,
                "collect_count": 0,
                "digg_count": 1,
                "comment_count": 0,
                "hot_index": 2,
                "is_hot": 0,
                "rank_index": 4e-8,
                "status": 1,
                "verify_status": 1,
                "audit_status": 2,
                "mark_content": "",
                "display_count": 0
            },
            "author_user_info": {
                "user_id": "123560414944158",
                "user_name": "Rust_Magazine",
                "company": "一本 Rust 语言开源杂志，每月末发刊",
                "job_title": "《Rust 中文精选》",
                "avatar_large": "https://p9-passport.byteacctimg.com/img/user-avatar/d0bd5f0426b8396a472f74eda0cf7808~300x300.image",
                "level": 3,
                "description": "定期更新 Rust 语言社区干货",
                "followee_count": 0,
                "follower_count": 463,
                "post_article_count": 198,
                "digg_article_count": 93,
                "got_digg_count": 557,
                "got_view_count": 162615,
                "post_shortmsg_count": 3,
                "digg_shortmsg_count": 1,
                "isfollowed": true,
                "favorable_author": 1,
                "power": 2185,
                "study_point": 0,
                "university": {
                    "university_id": "0",
                    "name": "",
                    "logo": ""
                },
                "major": {
                    "major_id": "0",
                    "parent_id": "0",
                    "name": ""
                },
                "student_status": 0,
                "select_event_count": 0,
                "select_online_course_count": 0,
                "identity": 0,
                "is_select_annual": false,
                "select_annual_rank": 0,
                "annual_list_type": 0,
                "extraMap": {},
                "is_logout": 0
            },
            "category": {
                "category_id": "6809637769959178254",
                "category_name": "后端",
                "category_url": "backend",
                "rank": 1,
                "back_ground": "https://lc-mhke0kuv.cn-n1.lcfile.com/fb3b208d06e6fe32.png",
                "icon": "https://lc-mhke0kuv.cn-n1.lcfile.com/a2ec01b816abd4c5.png",
                "ctime": 1457483880,
                "mtime": 1432503193,
                "show_type": 3,
                "item_type": 2,
                "promote_tag_cap": 4,
                "promote_priority": 1
            },
            "tags": [
                {
                    "id": 2546527,
                    "tag_id": "6809640408797167623",
                    "tag_name": "后端",
                    "color": "#C679FF",
                    "icon": "https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/leancloud-assets/d83da9d012ddb7ae85f4.png~tplv-t2oaga2asx-image.image",
                    "back_ground": "",
                    "show_navi": 1,
                    "ctime": 1435971556,
                    "mtime": 1641803746,
                    "id_type": 9,
                    "tag_alias": "",
                    "post_article_count": 95981,
                    "concern_user_count": 438297
                },
                {
                    "id": 2546590,
                    "tag_id": "6809640495594078216",
                    "tag_name": "Rust",
                    "color": "#616161",
                    "icon": "https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/leancloud-assets/01787a4f2859cde4db11.png~tplv-t2oaga2asx-image.image",
                    "back_ground": "",
                    "show_navi": 0,
                    "ctime": 1439498965,
                    "mtime": 1641801277,
                    "id_type": 9,
                    "tag_alias": "",
                    "post_article_count": 964,
                    "concern_user_count": 8662
                }
            ],
            "user_interact": {
                "id": 7050027158135112000,
                "omitempty": 2,
                "user_id": 3984285869016509,
                "is_digg": false,
                "is_follow": false,
                "is_collect": false
            },
            "org": {
                "org_info": null,
                "org_user": null,
                "is_followed": false
            },
            "req_id": "202201101640270102080650252144A87B",
            "status": {
                "push_status": 0
            }
        },
        {
            "article_id": "7050022764836552741",
            "article_info": {
                "article_id": "7050022764836552741",
                "user_id": "123560414944158",
                "category_id": "6809637769959178254",
                "tag_ids": [
                    6809640408797168000,
                    6809640495594078000
                ],
                "visible_level": 0,
                "link_url": "",
                "cover_image": "",
                "is_gfw": 0,
                "title": "Rust 社区 12 月热点",
                "brief_content": "聚焦 Rust 生态热点新闻 Rust for Linux 补丁更新到 V2 版本 2022 年，我们很可能会看到 Linux 内核中的实验性 Rust 编程语言支持成为主流。2021.12.6 早上",
                "is_english": 0,
                "is_original": 1,
                "user_index": 5.921972852168039,
                "original_type": 0,
                "original_author": "",
                "content": "",
                "ctime": "1641461361",
                "mtime": "1641467324",
                "rtime": "-62135596800",
                "draft_id": "7050021890827485191",
                "view_count": 24,
                "collect_count": 0,
                "digg_count": 0,
                "comment_count": 0,
                "hot_index": 1,
                "is_hot": 0,
                "rank_index": 3e-8,
                "status": 1,
                "verify_status": 1,
                "audit_status": 2,
                "mark_content": "",
                "display_count": 0
            },
            "author_user_info": {
                "user_id": "123560414944158",
                "user_name": "Rust_Magazine",
                "company": "一本 Rust 语言开源杂志，每月末发刊",
                "job_title": "《Rust 中文精选》",
                "avatar_large": "https://p9-passport.byteacctimg.com/img/user-avatar/d0bd5f0426b8396a472f74eda0cf7808~300x300.image",
                "level": 3,
                "description": "定期更新 Rust 语言社区干货",
                "followee_count": 0,
                "follower_count": 463,
                "post_article_count": 198,
                "digg_article_count": 93,
                "got_digg_count": 557,
                "got_view_count": 162615,
                "post_shortmsg_count": 3,
                "digg_shortmsg_count": 1,
                "isfollowed": true,
                "favorable_author": 1,
                "power": 2185,
                "study_point": 0,
                "university": {
                    "university_id": "0",
                    "name": "",
                    "logo": ""
                },
                "major": {
                    "major_id": "0",
                    "parent_id": "0",
                    "name": ""
                },
                "student_status": 0,
                "select_event_count": 0,
                "select_online_course_count": 0,
                "identity": 0,
                "is_select_annual": false,
                "select_annual_rank": 0,
                "annual_list_type": 0,
                "extraMap": {},
                "is_logout": 0
            },
            "category": {
                "category_id": "6809637769959178254",
                "category_name": "后端",
                "category_url": "backend",
                "rank": 1,
                "back_ground": "https://lc-mhke0kuv.cn-n1.lcfile.com/fb3b208d06e6fe32.png",
                "icon": "https://lc-mhke0kuv.cn-n1.lcfile.com/a2ec01b816abd4c5.png",
                "ctime": 1457483880,
                "mtime": 1432503193,
                "show_type": 3,
                "item_type": 2,
                "promote_tag_cap": 4,
                "promote_priority": 1
            },
            "tags": [
                {
                    "id": 2546527,
                    "tag_id": "6809640408797167623",
                    "tag_name": "后端",
                    "color": "#C679FF",
                    "icon": "https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/leancloud-assets/d83da9d012ddb7ae85f4.png~tplv-t2oaga2asx-image.image",
                    "back_ground": "",
                    "show_navi": 1,
                    "ctime": 1435971556,
                    "mtime": 1641803746,
                    "id_type": 9,
                    "tag_alias": "",
                    "post_article_count": 95981,
                    "concern_user_count": 438297
                },
                {
                    "id": 2546590,
                    "tag_id": "6809640495594078216",
                    "tag_name": "Rust",
                    "color": "#616161",
                    "icon": "https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/leancloud-assets/01787a4f2859cde4db11.png~tplv-t2oaga2asx-image.image",
                    "back_ground": "",
                    "show_navi": 0,
                    "ctime": 1439498965,
                    "mtime": 1641801277,
                    "id_type": 9,
                    "tag_alias": "",
                    "post_article_count": 964,
                    "concern_user_count": 8662
                }
            ],
            "user_interact": {
                "id": 7050022764836553000,
                "omitempty": 2,
                "user_id": 3984285869016509,
                "is_digg": false,
                "is_follow": false,
                "is_collect": false
            },
            "org": {
                "org_info": null,
                "org_user": null,
                "is_followed": false
            },
            "req_id": "202201101640270102080650252144A87B",
            "status": {
                "push_status": 0
            }
        },
        {
            "article_id": "7049990685113450503",
            "article_info": {
                "article_id": "7049990685113450503",
                "user_id": "4186596000416094",
                "category_id": "6809637769959178254",
                "tag_ids": [
                    6809640501776482000,
                    6809641127432421000
                ],
                "visible_level": 0,
                "link_url": "",
                "cover_image": "https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/8471eecdd2b549f19fe3ea69832482f6~tplv-k3u1fbpfcp-watermark.image?",
                "is_gfw": 0,
                "title": "深入理解百度在离线混部技术",
                "brief_content": "【百度云原生导读】服务器资源利用率较低，TCO(IT 基础设施的总拥有成本)逐年上涨，对于拥有大量机器资源的公司来说无疑是一个头疼的问题。混部技术就是在这种情况下应运而生，目前，混部技术在业界还属于比",
                "is_english": 0,
                "is_original": 1,
                "user_index": 9.083467274459665,
                "original_type": 0,
                "original_author": "",
                "content": "",
                "ctime": "1641453917",
                "mtime": "1641463982",
                "rtime": "-62135596800",
                "draft_id": "7049969091947266085",
                "view_count": 47,
                "collect_count": 1,
                "digg_count": 1,
                "comment_count": 0,
                "hot_index": 3,
                "is_hot": 0,
                "rank_index": 6e-8,
                "status": 1,
                "verify_status": 1,
                "audit_status": 2,
                "mark_content": "",
                "display_count": 0
            },
            "author_user_info": {
                "user_id": "4186596000416094",
                "user_name": "百度Geek说",
                "company": "百度",
                "job_title": "架构师",
                "avatar_large": "https://p9-passport.byteacctimg.com/img/user-avatar/722ef2acd6da963abceb16ce1f9b5cc3~300x300.image",
                "level": 3,
                "description": "公众号：百度Geek说",
                "followee_count": 5,
                "follower_count": 220,
                "post_article_count": 92,
                "digg_article_count": 92,
                "got_digg_count": 3111,
                "got_view_count": 65656,
                "post_shortmsg_count": 2,
                "digg_shortmsg_count": 1,
                "isfollowed": true,
                "favorable_author": 0,
                "power": 3772,
                "study_point": 0,
                "university": {
                    "university_id": "0",
                    "name": "",
                    "logo": ""
                },
                "major": {
                    "major_id": "0",
                    "parent_id": "0",
                    "name": ""
                },
                "student_status": 0,
                "select_event_count": 0,
                "select_online_course_count": 0,
                "identity": 0,
                "is_select_annual": true,
                "select_annual_rank": 0,
                "annual_list_type": 1,
                "extraMap": {},
                "is_logout": 0
            },
            "category": {
                "category_id": "6809637769959178254",
                "category_name": "后端",
                "category_url": "backend",
                "rank": 1,
                "back_ground": "https://lc-mhke0kuv.cn-n1.lcfile.com/fb3b208d06e6fe32.png",
                "icon": "https://lc-mhke0kuv.cn-n1.lcfile.com/a2ec01b816abd4c5.png",
                "ctime": 1457483880,
                "mtime": 1432503193,
                "show_type": 3,
                "item_type": 2,
                "promote_tag_cap": 4,
                "promote_priority": 1
            },
            "tags": [
                {
                    "id": 2546594,
                    "tag_id": "6809640501776482317",
                    "tag_name": "架构",
                    "color": "#C679FF",
                    "icon": "https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/leancloud-assets/f27d811ad7e2b2a0bc24.png~tplv-t2oaga2asx-image.image",
                    "back_ground": "",
                    "show_navi": 0,
                    "ctime": 1439515816,
                    "mtime": 1641803139,
                    "id_type": 9,
                    "tag_alias": "",
                    "post_article_count": 12458,
                    "concern_user_count": 349606
                },
                {
                    "id": 2547047,
                    "tag_id": "6809641127432421384",
                    "tag_name": "云原生",
                    "color": "",
                    "icon": "https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/gold-user-assets/1533870051840b17c1d74ad32d6b84756805faba0469c.jpg~tplv-t2oaga2asx-image.image",
                    "back_ground": "",
                    "show_navi": 0,
                    "ctime": 1533841257,
                    "mtime": 1641802500,
                    "id_type": 9,
                    "tag_alias": "Cloud Native",
                    "post_article_count": 1459,
                    "concern_user_count": 4254
                }
            ],
            "user_interact": {
                "id": 7049990685113450000,
                "omitempty": 2,
                "user_id": 3984285869016509,
                "is_digg": false,
                "is_follow": false,
                "is_collect": false
            },
            "org": {
                "org_info": {
                    "org_type": 1,
                    "org_id": "6933023521425588232",
                    "online_version_id": 6958058397480518000,
                    "latest_version_id": 6958058397480518000,
                    "power": 3639,
                    "ctime": 1614221648,
                    "mtime": 1641804018,
                    "audit_status": 2,
                    "status": 0,
                    "org_version": {
                        "version_id": "6958058397480517672",
                        "icon": "https://p9-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/c0a10aae4b104e94b105d27c16b608a5~tplv-k3u1fbpfcp-watermark.image",
                        "background": "https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/771d128a568f43ca8d37046a07415e34~tplv-k3u1fbpfcp-watermark.image",
                        "name": "百度Geek说",
                        "introduction": "关注我们，带你了解更多百度技术干货。\n\n",
                        "weibo_link": "",
                        "github_link": "",
                        "homepage_link": "",
                        "ctime": 1620269998,
                        "mtime": 1620269998,
                        "org_id": "6933023521425588232",
                        "brief_introduction": "【百度Geek说】公众号团队",
                        "introduction_preview": "关注我们，带你了解更多百度技术干货。"
                    },
                    "follower_count": 598,
                    "article_view_count": 64546,
                    "article_digg_count": 2994
                },
                "org_user": null,
                "is_followed": false
            },
            "req_id": "202201101640270102080650252144A87B",
            "status": {
                "push_status": 0
            }
        },
        {
            "article_id": "7049929787883651080",
            "article_info": {
                "article_id": "7049929787883651080",
                "user_id": "4186596000416094",
                "category_id": "6809637769959178254",
                "tag_ids": [
                    6809640501776482000,
                    6809640482725954000,
                    6809640408797168000
                ],
                "visible_level": 0,
                "link_url": "",
                "cover_image": "https://p1-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/ab8e05ff8884406b80f666a37e6035e7~tplv-k3u1fbpfcp-watermark.image?",
                "is_gfw": 0,
                "title": "当技术重构遇上DDD，如何实现业务、技术双赢？",
                "brief_content": "困境：项目背景 爱番番沟通基于百度商桥快速完成了产品功能和技术架构的从无到有，但同时也继承了百度商桥历史功能繁杂、技术架构陈旧的缺点。为了能更好地服务于爱番番",
                "is_english": 0,
                "is_original": 1,
                "user_index": 9.111024024410993,
                "original_type": 0,
                "original_author": "",
                "content": "",
                "ctime": "1641439732",
                "mtime": "1641538120",
                "rtime": "1641538120",
                "draft_id": "7049916982337994759",
                "view_count": 2696,
                "collect_count": 26,
                "digg_count": 21,
                "comment_count": 0,
                "hot_index": 155,
                "is_hot": 0,
                "rank_index": 0.60109403,
                "status": 2,
                "verify_status": 1,
                "audit_status": 2,
                "mark_content": "",
                "display_count": 0
            },
            "author_user_info": {
                "user_id": "4186596000416094",
                "user_name": "百度Geek说",
                "company": "百度",
                "job_title": "架构师",
                "avatar_large": "https://p9-passport.byteacctimg.com/img/user-avatar/722ef2acd6da963abceb16ce1f9b5cc3~300x300.image",
                "level": 3,
                "description": "公众号：百度Geek说",
                "followee_count": 5,
                "follower_count": 220,
                "post_article_count": 92,
                "digg_article_count": 92,
                "got_digg_count": 3111,
                "got_view_count": 65656,
                "post_shortmsg_count": 2,
                "digg_shortmsg_count": 1,
                "isfollowed": true,
                "favorable_author": 0,
                "power": 3772,
                "study_point": 0,
                "university": {
                    "university_id": "0",
                    "name": "",
                    "logo": ""
                },
                "major": {
                    "major_id": "0",
                    "parent_id": "0",
                    "name": ""
                },
                "student_status": 0,
                "select_event_count": 0,
                "select_online_course_count": 0,
                "identity": 0,
                "is_select_annual": true,
                "select_annual_rank": 0,
                "annual_list_type": 1,
                "extraMap": {},
                "is_logout": 0
            },
            "category": {
                "category_id": "6809637769959178254",
                "category_name": "后端",
                "category_url": "backend",
                "rank": 1,
                "back_ground": "https://lc-mhke0kuv.cn-n1.lcfile.com/fb3b208d06e6fe32.png",
                "icon": "https://lc-mhke0kuv.cn-n1.lcfile.com/a2ec01b816abd4c5.png",
                "ctime": 1457483880,
                "mtime": 1432503193,
                "show_type": 3,
                "item_type": 2,
                "promote_tag_cap": 4,
                "promote_priority": 1
            },
            "tags": [
                {
                    "id": 2546594,
                    "tag_id": "6809640501776482317",
                    "tag_name": "架构",
                    "color": "#C679FF",
                    "icon": "https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/leancloud-assets/f27d811ad7e2b2a0bc24.png~tplv-t2oaga2asx-image.image",
                    "back_ground": "",
                    "show_navi": 0,
                    "ctime": 1439515816,
                    "mtime": 1641803139,
                    "id_type": 9,
                    "tag_alias": "",
                    "post_article_count": 12458,
                    "concern_user_count": 349606
                },
                {
                    "id": 2546581,
                    "tag_id": "6809640482725953550",
                    "tag_name": "程序员",
                    "color": "#616161",
                    "icon": "https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/leancloud-assets/63baec1130bde0284e98.png~tplv-t2oaga2asx-image.image",
                    "back_ground": "",
                    "show_navi": 0,
                    "ctime": 1438712834,
                    "mtime": 1641803557,
                    "id_type": 9,
                    "tag_alias": "",
                    "post_article_count": 18276,
                    "concern_user_count": 284059
                },
                {
                    "id": 2546527,
                    "tag_id": "6809640408797167623",
                    "tag_name": "后端",
                    "color": "#C679FF",
                    "icon": "https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/leancloud-assets/d83da9d012ddb7ae85f4.png~tplv-t2oaga2asx-image.image",
                    "back_ground": "",
                    "show_navi": 1,
                    "ctime": 1435971556,
                    "mtime": 1641803746,
                    "id_type": 9,
                    "tag_alias": "",
                    "post_article_count": 95981,
                    "concern_user_count": 438297
                }
            ],
            "user_interact": {
                "id": 7049929787883651000,
                "omitempty": 2,
                "user_id": 3984285869016509,
                "is_digg": false,
                "is_follow": false,
                "is_collect": false
            },
            "org": {
                "org_info": {
                    "org_type": 1,
                    "org_id": "6933023521425588232",
                    "online_version_id": 6958058397480518000,
                    "latest_version_id": 6958058397480518000,
                    "power": 3639,
                    "ctime": 1614221648,
                    "mtime": 1641804018,
                    "audit_status": 2,
                    "status": 0,
                    "org_version": {
                        "version_id": "6958058397480517672",
                        "icon": "https://p9-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/c0a10aae4b104e94b105d27c16b608a5~tplv-k3u1fbpfcp-watermark.image",
                        "background": "https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/771d128a568f43ca8d37046a07415e34~tplv-k3u1fbpfcp-watermark.image",
                        "name": "百度Geek说",
                        "introduction": "关注我们，带你了解更多百度技术干货。\n\n",
                        "weibo_link": "",
                        "github_link": "",
                        "homepage_link": "",
                        "ctime": 1620269998,
                        "mtime": 1620269998,
                        "org_id": "6933023521425588232",
                        "brief_introduction": "【百度Geek说】公众号团队",
                        "introduction_preview": "关注我们，带你了解更多百度技术干货。"
                    },
                    "follower_count": 598,
                    "article_view_count": 64546,
                    "article_digg_count": 2994
                },
                "org_user": null,
                "is_followed": false
            },
            "req_id": "202201101640270102080650252144A87B",
            "status": {
                "push_status": 0
            }
        },
        {
            "article_id": "7049886535339999239",
            "article_info": {
                "article_id": "7049886535339999239",
                "user_id": "1574156384091320",
                "category_id": "6809637771511070734",
                "tag_ids": [
                    6809640375880253000,
                    6809640419505209000
                ],
                "visible_level": 0,
                "link_url": "",
                "cover_image": "https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/ab4f1c4943f34f9594a1083bfc8fe85a~tplv-k3u1fbpfcp-watermark.image?",
                "is_gfw": 0,
                "title": "在 GitHub 复活 80 年代的游戏代码，它们出自第一本售出百万册的计算机书籍",
                "brief_content": "今儿我在 GitHub 看到了一个很眼熟的名字和头像，但是第一时间没想起来他是谁。算了先看看是个什么神仙开源项目，竟然能登上今天的 GitHub 趋势榜首。 该项目是把《BASIC Computer ",
                "is_english": 0,
                "is_original": 1,
                "user_index": 9.309165713609845,
                "original_type": 0,
                "original_author": "",
                "content": "",
                "ctime": "1641429692",
                "mtime": "1641537897",
                "rtime": "1641537897",
                "draft_id": "7049741613634945061",
                "view_count": 1639,
                "collect_count": 0,
                "digg_count": 3,
                "comment_count": 0,
                "hot_index": 84,
                "is_hot": 0,
                "rank_index": 0.34139984,
                "status": 2,
                "verify_status": 1,
                "audit_status": 2,
                "mark_content": "",
                "display_count": 0
            },
            "author_user_info": {
                "user_id": "1574156384091320",
                "user_name": "HelloGitHub",
                "company": "公众号：HelloGitHub",
                "job_title": "",
                "avatar_large": "https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/gold-user-assets/2019/3/28/169c31b06dcd6e90~tplv-t2oaga2asx-image.image",
                "level": 5,
                "description": "找开源项目可以试试 HelloGitHub 小程序～",
                "followee_count": 33,
                "follower_count": 10802,
                "post_article_count": 376,
                "digg_article_count": 264,
                "got_digg_count": 9118,
                "got_view_count": 810970,
                "post_shortmsg_count": 1087,
                "digg_shortmsg_count": 220,
                "isfollowed": true,
                "favorable_author": 1,
                "power": 16363,
                "study_point": 0,
                "university": {
                    "university_id": "0",
                    "name": "",
                    "logo": ""
                },
                "major": {
                    "major_id": "0",
                    "parent_id": "0",
                    "name": ""
                },
                "student_status": 0,
                "select_event_count": 0,
                "select_online_course_count": 0,
                "identity": 0,
                "is_select_annual": true,
                "select_annual_rank": 0,
                "annual_list_type": 0,
                "extraMap": {},
                "is_logout": 0
            },
            "category": {
                "category_id": "6809637771511070734",
                "category_name": "开发工具",
                "category_url": "freebie",
                "rank": 6,
                "back_ground": "https://lc-mhke0kuv.cn-n1.lcfile.com/412957a61f414c0b.png",
                "icon": "https://lc-mhke0kuv.cn-n1.lcfile.com/eb20ab1334d9abea.png",
                "ctime": 1457483920,
                "mtime": 1432503202,
                "show_type": 3,
                "item_type": 2,
                "promote_tag_cap": 4,
                "promote_priority": 6
            },
            "tags": [
                {
                    "id": 2546502,
                    "tag_id": "6809640375880253447",
                    "tag_name": "GitHub",
                    "color": "#616161",
                    "icon": "https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/leancloud-assets/0d614af263aa63aa6a77.png~tplv-t2oaga2asx-image.image",
                    "back_ground": "",
                    "show_navi": 0,
                    "ctime": 1432234558,
                    "mtime": 1641803412,
                    "id_type": 9,
                    "tag_alias": "",
                    "post_article_count": 11281,
                    "concern_user_count": 397818
                },
                {
                    "id": 2546535,
                    "tag_id": "6809640419505209358",
                    "tag_name": "开源",
                    "color": "#6EBD68",
                    "icon": "https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/leancloud-assets/553ecacd498946a9a6d9.png~tplv-t2oaga2asx-image.image",
                    "back_ground": "",
                    "show_navi": 0,
                    "ctime": 1435972427,
                    "mtime": 1641803418,
                    "id_type": 9,
                    "tag_alias": "",
                    "post_article_count": 6377,
                    "concern_user_count": 220213
                }
            ],
            "user_interact": {
                "id": 7049886535339999000,
                "omitempty": 2,
                "user_id": 3984285869016509,
                "is_digg": false,
                "is_follow": false,
                "is_collect": false
            },
            "org": {
                "org_info": null,
                "org_user": null,
                "is_followed": false
            },
            "req_id": "202201101640270102080650252144A87B",
            "status": {
                "push_status": 0
            }
        },
        {
            "article_id": "7049709896983379981",
            "article_info": {
                "article_id": "7049709896983379981",
                "user_id": "1521379821230269",
                "category_id": "6809637769959178254",
                "tag_ids": [
                    6809640448827589000,
                    6809640415537398000,
                    6809640408797168000
                ],
                "visible_level": 0,
                "link_url": "",
                "cover_image": "https://p1-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/35f6712b9c5346098f8af6dc0cab0f83~tplv-k3u1fbpfcp-watermark.image?",
                "is_gfw": 0,
                "title": "两行代码，给 Python 脚本生成命令行",
                "brief_content": "有时候我们会有这样的一个需求： 我们定义了一个 Python 的方法，方法接收一些参数，但是调用的时候想将这些参数用命令行暴露出来。 比如说这里有个爬取方法： 这里定义了一个 scrape 方法，第一",
                "is_english": 0,
                "is_original": 1,
                "user_index": 11.215727466383392,
                "original_type": 0,
                "original_author": "",
                "content": "",
                "ctime": "1641388683",
                "mtime": "1641455918",
                "rtime": "1641455918",
                "draft_id": "7049709973542830088",
                "view_count": 1700,
                "collect_count": 4,
                "digg_count": 4,
                "comment_count": 0,
                "hot_index": 89,
                "is_hot": 0,
                "rank_index": 0.25963339,
                "status": 2,
                "verify_status": 1,
                "audit_status": 2,
                "mark_content": "",
                "display_count": 0
            },
            "author_user_info": {
                "user_id": "1521379821230269",
                "user_name": "崔庆才丨静觅",
                "company": "微软（中国）有限公司",
                "job_title": "工程师",
                "avatar_large": "https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/gold-user-assets/2018/3/8/1620484dfb894437~tplv-t2oaga2asx-image.image",
                "level": 4,
                "description": "静静寻觅生活的美好。",
                "followee_count": 21,
                "follower_count": 9860,
                "post_article_count": 95,
                "digg_article_count": 56,
                "got_digg_count": 4380,
                "got_view_count": 424073,
                "post_shortmsg_count": 1,
                "digg_shortmsg_count": 4,
                "isfollowed": true,
                "favorable_author": 1,
                "power": 8934,
                "study_point": 0,
                "university": {
                    "university_id": "0",
                    "name": "",
                    "logo": ""
                },
                "major": {
                    "major_id": "0",
                    "parent_id": "0",
                    "name": ""
                },
                "student_status": 0,
                "select_event_count": 0,
                "select_online_course_count": 0,
                "identity": 0,
                "is_select_annual": false,
                "select_annual_rank": 0,
                "annual_list_type": 0,
                "extraMap": {},
                "is_logout": 0
            },
            "category": {
                "category_id": "6809637769959178254",
                "category_name": "后端",
                "category_url": "backend",
                "rank": 1,
                "back_ground": "https://lc-mhke0kuv.cn-n1.lcfile.com/fb3b208d06e6fe32.png",
                "icon": "https://lc-mhke0kuv.cn-n1.lcfile.com/a2ec01b816abd4c5.png",
                "ctime": 1457483880,
                "mtime": 1432503193,
                "show_type": 3,
                "item_type": 2,
                "promote_tag_cap": 4,
                "promote_priority": 1
            },
            "tags": [
                {
                    "id": 2546556,
                    "tag_id": "6809640448827588622",
                    "tag_name": "Python",
                    "color": "#356E9C",
                    "icon": "https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/leancloud-assets/b51a1dacf9bb7883defe.png~tplv-t2oaga2asx-image.image",
                    "back_ground": "",
                    "show_navi": 0,
                    "ctime": 1436156327,
                    "mtime": 1641803592,
                    "id_type": 9,
                    "tag_alias": "",
                    "post_article_count": 29018,
                    "concern_user_count": 221630
                },
                {
                    "id": 2546532,
                    "tag_id": "6809640415537397768",
                    "tag_name": "命令行",
                    "color": "#F56868",
                    "icon": "https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/leancloud-assets/03f4a3a6b66a8a91627d.png~tplv-t2oaga2asx-image.image",
                    "back_ground": "",
                    "show_navi": 0,
                    "ctime": 1435972175,
                    "mtime": 1641795144,
                    "id_type": 9,
                    "tag_alias": "",
                    "post_article_count": 3459,
                    "concern_user_count": 119438
                },
                {
                    "id": 2546527,
                    "tag_id": "6809640408797167623",
                    "tag_name": "后端",
                    "color": "#C679FF",
                    "icon": "https://p1-jj.byteimg.com/tos-cn-i-t2oaga2asx/leancloud-assets/d83da9d012ddb7ae85f4.png~tplv-t2oaga2asx-image.image",
                    "back_ground": "",
                    "show_navi": 1,
                    "ctime": 1435971556,
                    "mtime": 1641803746,
                    "id_type": 9,
                    "tag_alias": "",
                    "post_article_count": 95981,
                    "concern_user_count": 438297
                }
            ],
            "user_interact": {
                "id": 7049709896983380000,
                "omitempty": 2,
                "user_id": 3984285869016509,
                "is_digg": false,
                "is_follow": false,
                "is_collect": false
            },
            "org": {
                "org_info": null,
                "org_user": null,
                "is_followed": false
            },
            "req_id": "202201101640270102080650252144A87B",
            "status": {
                "push_status": 0
            }
        }
    ],
    "cursor": "10",
    "count": 0,
    "has_more": true
}
`

var webData = `
{"ok":1,"statuses":[{"visible":{"type":0,"list_id":0},"created_at":"Wed Mar 23 16:18:20 +0800 2022","id":4750217209449536,"idstr":"4750217209449536","mid":"4750217209449536","mblogid":"Ll4nuDEfS","user":{"id":3612749480,"idstr":"3612749480","pc_new":6,"screen_name":"山人I","profile_image_url":"https://tvax3.sinaimg.cn/crop.0.0.828.828.50/d7562ea8ly8gnyu8bneyoj20n00n0jt8.jpg?KID=imgbed,tva&Expires=1648035252&ssig=0jrxHNyYN6","profile_url":"/u/3612749480","verified":true,"verified_type":0,"domain":"","weihao":"","verified_type_ext":1,"avatar_large":"https://tvax3.sinaimg.cn/crop.0.0.828.828.180/d7562ea8ly8gnyu8bneyoj20n00n0jt8.jpg?KID=imgbed,tva&Expires=1648035252&ssig=Lq%2BaZDZh9G","avatar_hd":"https://tvax3.sinaimg.cn/crop.0.0.828.828.1024/d7562ea8ly8gnyu8bneyoj20n00n0jt8.jpg?KID=imgbed,tva&Expires=1648035252&ssig=YXtkeuJPA5","follow_me":false,"following":true,"mbrank":7,"mbtype":12,"planet_video":false,"icon_list":[{"type":"vip","data":{"mbrank":7,"mbtype":12}}]},"can_edit":false,"text_raw":"反弹六天了，指标炒买，马上就要拐头了。\n缩量这样涨一点意义都没有，反而不是好事。\n一旦破缺口麻烦就大了，按道理会拉高上去一个平台再回撤一次，但明显弱势，要是大起大落还好，拉银行地产很费资金，一旦它们开始回撤如果不能轮动起来那就被带坑里去了。\n就怕美股跌了，直接就被吓死，加上北资做空 ​​​","text":"反弹六天了，指标炒买，马上就要拐头了。<br />缩量这样涨一点意义都没有，反而不是好事。<br />一旦破缺口麻烦就大了，按道理会拉高上去一个平台再回撤一次，但明显弱势，要是大起大落还好，拉银行地产很费资金，一旦它们开始回撤如果不能轮动起来那就被带坑里去了。<br />就怕美股跌了，直接就被吓死，加上北资做空 ​​​ ...<span class=\"expand\">展开</span>","textLength":402,"source":"iPhone 11","favorited":false,"rid":"0_0_200_6665576152895923232_0_0_0","is_controlled_by_server":"0","pic_ids":[],"geo":null,"pic_num":0,"is_paid":false,"mblog_vip_type":0,"number_display_strategy":{"apply_scenario_flag":3,"display_text_min_number":1000000,"display_text":"100万+"},"reposts_count":2,"comments_count":176,"attitudes_count":348,"attitudes_status":0,"continue_tag":{"title":"全文","pic":"http://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article.png","scheme":"sinaweibo://detail?mblogid=4750217209449536&id=4750217209449536"},"isLongText":true,"mlevel":0,"content_auth":0,"is_show_bulletin":2,"comment_manage_info":{"comment_permission_type":-1,"approval_comment_type":0,"comment_sort_type":0},"share_repost_type":0,"tag_struct":[{"tag_name":"纳斯达克 14108.82","oid":"1022:230677.IXIC","tag_type":2,"tag_hidden":0,"tag_scheme":"https://stock.weibo.cn/page/tag?oid=1022:230677.IXIC","url_type_pic":"https://h5.sinaimg.cn/upload/1008/253/2018/12/26/timeline_icon_stock.png","actionlog":{"act_code":2413,"oid":"1022:230677.IXIC","uicode":null,"luicode":null,"fid":null,"ext":"|tag_type:stock"},"bd_object_type":"stock","desc":"+1.95%"}],"mblogtype":0,"showFeedRepost":false,"showFeedComment":false,"pictureViewerSign":false,"showPictureViewer":false,"rcList":[],"customIcons":[]},{"visible":{"type":0,"list_id":0},"created_at":"Wed Mar 23 16:09:04 +0800 2022","id":4750214877416477,"idstr":"4750214877416477","mid":"4750214877416477","mblogid":"Ll4jJv7mB","user":{"id":1402070602,"idstr":"1402070602","pc_new":7,"screen_name":"诺饭短线","profile_image_url":"https://tva1.sinaimg.cn/crop.0.0.180.180.50/5391e64ajw1e8qgp5bmzyj2050050aa8.jpg?KID=imgbed,tva&Expires=1648035252&ssig=NpmsWD8KPj","profile_url":"/u/1402070602","verified":true,"verified_type":0,"domain":"swgreat","weihao":"275745544","verified_type_ext":1,"avatar_large":"https://tva1.sinaimg.cn/crop.0.0.180.180.180/5391e64ajw1e8qgp5bmzyj2050050aa8.jpg?KID=imgbed,tva&Expires=1648035252&ssig=QrIREWj3n0","avatar_hd":"https://tva1.sinaimg.cn/crop.0.0.180.180.1024/5391e64ajw1e8qgp5bmzyj2050050aa8.jpg?KID=imgbed,tva&Expires=1648035252&ssig=SifMXru1Ea","follow_me":false,"following":true,"mbrank":6,"mbtype":11,"planet_video":true,"icon_list":[{"type":"vip","data":{"mbrank":6,"mbtype":11}}]},"can_edit":false,"text_raw":"昨天才上宋都，不管是打板还是低吸，今天跌停不走，而是等到平盘附近走的，绝对是高手，做短线是很难亏钱的。我肯定是做不到这样的。即使我知道7板高度只要能封住，一般来说是不会让打板族亏钱的的，也很可能会因为恐慌或者失望情绪导致提前卖出的。 \n实际上你可以看一下春节至今的7板股，没有一个是7 ​​​","text":"昨天才上宋都，不管是打板还是低吸，今天跌停不走，而是等到平盘附近走的，绝对是高手，做短线是很难亏钱的。我肯定是做不到这样的。即使我知道7板高度只要能封住，一般来说是不会让打板族亏钱的的，也很可能会因为恐慌或者失望情绪导致提前卖出的。 <br />实际上你可以看一下春节至今的7板股，没有一个是7 ​​​ ...<span class=\"expand\">展开</span>","textLength":371,"source":"新版微博 weibo.com","favorited":false,"rid":"1_0_200_6665576152895923232_0_0_0","is_controlled_by_server":"0","pic_ids":[],"geo":null,"pic_num":0,"is_paid":false,"mblog_vip_type":0,"number_display_strategy":{"apply_scenario_flag":3,"display_text_min_number":1000000,"display_text":"100万+"},"reposts_count":1,"comments_count":7,"attitudes_count":17,"attitudes_status":0,"continue_tag":{"title":"全文","pic":"http://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article.png","scheme":"sinaweibo://detail?mblogid=4750214877416477&id=4750214877416477"},"isLongText":true,"mlevel":0,"content_auth":0,"is_show_bulletin":2,"comment_manage_info":{"comment_permission_type":-1,"approval_comment_type":0,"comment_sort_type":0},"share_repost_type":0,"mblogtype":0,"showFeedRepost":false,"showFeedComment":false,"pictureViewerSign":false,"showPictureViewer":false,"rcList":[],"customIcons":[]},{"visible":{"type":0,"list_id":0},"created_at":"Wed Mar 23 15:55:38 +0800 2022","id":4750211497330341,"idstr":"4750211497330341","mid":"4750211497330341","mblogid":"Ll4ehuKXj","user":{"id":1402070602,"idstr":"1402070602","pc_new":7,"screen_name":"诺饭短线","profile_image_url":"https://tva1.sinaimg.cn/crop.0.0.180.180.50/5391e64ajw1e8qgp5bmzyj2050050aa8.jpg?KID=imgbed,tva&Expires=1648035252&ssig=NpmsWD8KPj","profile_url":"/u/1402070602","verified":true,"verified_type":0,"domain":"swgreat","weihao":"275745544","verified_type_ext":1,"avatar_large":"https://tva1.sinaimg.cn/crop.0.0.180.180.180/5391e64ajw1e8qgp5bmzyj2050050aa8.jpg?KID=imgbed,tva&Expires=1648035252&ssig=QrIREWj3n0","avatar_hd":"https://tva1.sinaimg.cn/crop.0.0.180.180.1024/5391e64ajw1e8qgp5bmzyj2050050aa8.jpg?KID=imgbed,tva&Expires=1648035252&ssig=SifMXru1Ea","follow_me":false,"following":true,"mbrank":6,"mbtype":11,"planet_video":true,"icon_list":[{"type":"vip","data":{"mbrank":6,"mbtype":11}}]},"can_edit":false,"text_raw":"房地产下午回流的时候，估计会有人去做宋都和苏州高新的反拉。其实这俩是最不应该做的。如果能起来，上午的分时就不会是这样的了。现在高标基本没有什么地天板之类的了，强就是强弱就是弱。最多你就可以打打荣安，后面的就都是卖出为主了。当然苏州高新我确实没想到能拉这么高，反拉均线附近就走了。但 ​​​","text":"房地产下午回流的时候，估计会有人去做宋都和苏州高新的反拉。其实这俩是最不应该做的。如果能起来，上午的分时就不会是这样的了。现在高标基本没有什么地天板之类的了，强就是强弱就是弱。最多你就可以打打荣安，后面的就都是卖出为主了。当然苏州高新我确实没想到能拉这么高，反拉均线附近就走了。但 ​​​ ...<span class=\"expand\">展开</span>","textLength":760,"source":"新版微博 weibo.com","favorited":false,"rid":"2_0_200_6665576152895923232_0_0_0","is_controlled_by_server":"0","pic_ids":[],"geo":null,"pic_num":0,"is_paid":false,"mblog_vip_type":0,"number_display_strategy":{"apply_scenario_flag":3,"display_text_min_number":1000000,"display_text":"100万+"},"reposts_count":1,"comments_count":16,"attitudes_count":18,"attitudes_status":0,"continue_tag":{"title":"全文","pic":"http://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article.png","scheme":"sinaweibo://detail?mblogid=4750211497330341&id=4750211497330341"},"isLongText":true,"mlevel":0,"content_auth":0,"is_show_bulletin":2,"comment_manage_info":{"comment_permission_type":-1,"approval_comment_type":0,"comment_sort_type":0},"share_repost_type":0,"tag_struct":[{"tag_name":"苏州高新 sh600736","oid":"1022:230677sh600736","tag_type":2,"tag_hidden":0,"tag_scheme":"https://stock.weibo.cn/page/tag?oid=1022:230677sh600736","url_type_pic":"https://h5.sinaimg.cn/upload/1008/253/2018/12/26/timeline_icon_stock.png","actionlog":{"act_code":2413,"oid":"1022:230677sh600736","uicode":null,"luicode":null,"fid":null,"ext":"|tag_type:stock"},"bd_object_type":"stock","desc":"-8.69%"}],"mblogtype":0,"showFeedRepost":false,"showFeedComment":false,"pictureViewerSign":false,"showPictureViewer":false,"rcList":[],"customIcons":[]},{"visible":{"type":0,"list_id":0},"created_at":"Wed Mar 23 15:35:43 +0800 2022","id":4750206484616141,"idstr":"4750206484616141","mid":"4750206484616141","mblogid":"Ll46cjmRT","user":{"id":3612749480,"idstr":"3612749480","pc_new":6,"screen_name":"山人I","profile_image_url":"https://tvax3.sinaimg.cn/crop.0.0.828.828.50/d7562ea8ly8gnyu8bneyoj20n00n0jt8.jpg?KID=imgbed,tva&Expires=1648035252&ssig=0jrxHNyYN6","profile_url":"/u/3612749480","verified":true,"verified_type":0,"domain":"","weihao":"","verified_type_ext":1,"avatar_large":"https://tvax3.sinaimg.cn/crop.0.0.828.828.180/d7562ea8ly8gnyu8bneyoj20n00n0jt8.jpg?KID=imgbed,tva&Expires=1648035252&ssig=Lq%2BaZDZh9G","avatar_hd":"https://tvax3.sinaimg.cn/crop.0.0.828.828.1024/d7562ea8ly8gnyu8bneyoj20n00n0jt8.jpg?KID=imgbed,tva&Expires=1648035252&ssig=YXtkeuJPA5","follow_me":false,"following":true,"mbrank":7,"mbtype":12,"planet_video":false,"icon_list":[{"type":"vip","data":{"mbrank":7,"mbtype":12}}]},"can_edit":false,"text_raw":"清明节马上到了，这两天估计要减仓了，这走法不对劲。\n拉高太多人卖，机构不出力没法子的。我考虑过无数次等拉个大阳线走人，下月再回来。\n蓝筹轮着拉都起不来，一旦回撤就真傻眼了。还要探底。这几天是美股和港股大涨硬顶着，就这样都不反弹，说明市场人气越来越低迷。 ​​​","text":"清明节马上到了，这两天估计要减仓了，这走法不对劲。<br />拉高太多人卖，机构不出力没法子的。我考虑过无数次等拉个大阳线走人，下月再回来。<br />蓝筹轮着拉都起不来，一旦回撤就真傻眼了。还要探底。这几天是美股和港股大涨硬顶着，就这样都不反弹，说明市场人气越来越低迷。 ​​​","textLength":250,"source":"iPhone 11","favorited":false,"rid":"3_0_200_6665576152895923232_0_0_0","is_controlled_by_server":"0","pic_ids":[],"geo":null,"pic_num":0,"is_paid":false,"mblog_vip_type":0,"number_display_strategy":{"apply_scenario_flag":3,"display_text_min_number":1000000,"display_text":"100万+"},"reposts_count":5,"comments_count":254,"attitudes_count":879,"attitudes_status":0,"isLongText":false,"mlevel":0,"content_auth":0,"is_show_bulletin":2,"comment_manage_info":{"comment_permission_type":-1,"approval_comment_type":0,"comment_sort_type":0},"share_repost_type":0,"mblogtype":0,"showFeedRepost":false,"showFeedComment":false,"pictureViewerSign":false,"showPictureViewer":false,"rcList":[],"customIcons":[]},{"visible":{"type":0,"list_id":0},"created_at":"Wed Mar 23 15:29:47 +0800 2022","id":4750204991702562,"idstr":"4750204991702562","mid":"4750204991702562","mblogid":"Ll43N78UG","user":{"id":1402070602,"idstr":"1402070602","pc_new":7,"screen_name":"诺饭短线","profile_image_url":"https://tva1.sinaimg.cn/crop.0.0.180.180.50/5391e64ajw1e8qgp5bmzyj2050050aa8.jpg?KID=imgbed,tva&Expires=1648035252&ssig=NpmsWD8KPj","profile_url":"/u/1402070602","verified":true,"verified_type":0,"domain":"swgreat","weihao":"275745544","verified_type_ext":1,"avatar_large":"https://tva1.sinaimg.cn/crop.0.0.180.180.180/5391e64ajw1e8qgp5bmzyj2050050aa8.jpg?KID=imgbed,tva&Expires=1648035252&ssig=QrIREWj3n0","avatar_hd":"https://tva1.sinaimg.cn/crop.0.0.180.180.1024/5391e64ajw1e8qgp5bmzyj2050050aa8.jpg?KID=imgbed,tva&Expires=1648035252&ssig=SifMXru1Ea","follow_me":false,"following":true,"mbrank":6,"mbtype":11,"planet_video":true,"icon_list":[{"type":"vip","data":{"mbrank":6,"mbtype":11}}]},"can_edit":false,"text_raw":"低吸策略一般适用的条件是，一，牛市，现在是熊市，所以整体不适用。二，上升期。现在总龙已经见顶了，所以是下降期。三，单日情绪由极弱转强。今天显然并不是转强。所以低吸都不适合。那明天是否能低吸搏长腿呢，标的是啥呢？能不能转强我不知道。不过标的容易选，就是今日烂板。又有一定辨识度的，并 ​​​","text":"低吸策略一般适用的条件是，一，牛市，现在是熊市，所以整体不适用。二，上升期。现在总龙已经见顶了，所以是下降期。三，单日情绪由极弱转强。今天显然并不是转强。所以低吸都不适合。那明天是否能低吸搏长腿呢，标的是啥呢？能不能转强我不知道。不过标的容易选，就是今日烂板。又有一定辨识度的，并 ​​​ ...<span class=\"expand\">展开</span>","textLength":438,"source":"新版微博 weibo.com","favorited":false,"rid":"4_0_200_6665576152895923232_0_0_0","is_controlled_by_server":"0","pic_ids":[],"geo":null,"pic_num":0,"is_paid":false,"mblog_vip_type":0,"number_display_strategy":{"apply_scenario_flag":3,"display_text_min_number":1000000,"display_text":"100万+"},"reposts_count":5,"comments_count":11,"attitudes_count":25,"attitudes_status":0,"continue_tag":{"title":"全文","pic":"http://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article.png","scheme":"sinaweibo://detail?mblogid=4750204991702562&id=4750204991702562"},"isLongText":true,"mlevel":0,"content_auth":0,"is_show_bulletin":2,"comment_manage_info":{"comment_permission_type":-1,"approval_comment_type":0,"comment_sort_type":0},"share_repost_type":0,"tag_struct":[{"tag_name":"宁波建工 sh601789","oid":"1022:230677sh601789","tag_type":2,"tag_hidden":0,"tag_scheme":"https://stock.weibo.cn/page/tag?oid=1022:230677sh601789","url_type_pic":"https://h5.sinaimg.cn/upload/1008/253/2018/12/26/timeline_icon_stock.png","actionlog":{"act_code":2413,"oid":"1022:230677sh601789","uicode":null,"luicode":null,"fid":null,"ext":"|tag_type:stock"},"bd_object_type":"stock","desc":"+10.01%"}],"mblogtype":0,"showFeedRepost":false,"showFeedComment":false,"pictureViewerSign":false,"showPictureViewer":false,"rcList":[],"customIcons":[]},{"visible":{"type":0,"list_id":0},"created_at":"Wed Mar 23 15:26:21 +0800 2022","id":4750204127939401,"idstr":"4750204127939401","mid":"4750204127939401","mblogid":"Ll42oxjoR","user":{"id":3612749480,"idstr":"3612749480","pc_new":6,"screen_name":"山人I","profile_image_url":"https://tvax3.sinaimg.cn/crop.0.0.828.828.50/d7562ea8ly8gnyu8bneyoj20n00n0jt8.jpg?KID=imgbed,tva&Expires=1648035252&ssig=0jrxHNyYN6","profile_url":"/u/3612749480","verified":true,"verified_type":0,"domain":"","weihao":"","verified_type_ext":1,"avatar_large":"https://tvax3.sinaimg.cn/crop.0.0.828.828.180/d7562ea8ly8gnyu8bneyoj20n00n0jt8.jpg?KID=imgbed,tva&Expires=1648035252&ssig=Lq%2BaZDZh9G","avatar_hd":"https://tvax3.sinaimg.cn/crop.0.0.828.828.1024/d7562ea8ly8gnyu8bneyoj20n00n0jt8.jpg?KID=imgbed,tva&Expires=1648035252&ssig=YXtkeuJPA5","follow_me":false,"following":true,"mbrank":7,"mbtype":12,"planet_video":false,"icon_list":[{"type":"vip","data":{"mbrank":7,"mbtype":12}}]},"can_edit":false,"text_raw":"搞笑的是券商一天波动不到1点，踢来踢去精准的很，结果一算扣了手续费剩下250，真是法克密了。\n哎，我快要无聊了死了！ ​​​","text":"搞笑的是券商一天波动不到1点，踢来踢去精准的很，结果一算扣了手续费剩下250，真是法克密了。<br />哎，我快要无聊了死了！ ​​​","textLength":111,"source":"iPhone 11","favorited":false,"rid":"5_0_200_6665576152895923232_0_0_0","is_controlled_by_server":"0","pic_ids":[],"geo":null,"pic_num":0,"is_paid":false,"mblog_vip_type":0,"number_display_strategy":{"apply_scenario_flag":3,"display_text_min_number":1000000,"display_text":"100万+"},"reposts_count":2,"comments_count":171,"attitudes_count":596,"attitudes_status":0,"isLongText":false,"mlevel":0,"content_auth":0,"is_show_bulletin":2,"comment_manage_info":{"comment_permission_type":-1,"approval_comment_type":0,"comment_sort_type":0},"share_repost_type":0,"mblogtype":0,"showFeedRepost":false,"showFeedComment":false,"pictureViewerSign":false,"showPictureViewer":false,"rcList":[],"customIcons":[]},{"visible":{"type":0,"list_id":0},"created_at":"Wed Mar 23 15:25:35 +0800 2022","id":4750203934215156,"idstr":"4750203934215156","mid":"4750203934215156","mblogid":"Ll425hGyo","user":{"id":1196703900,"idstr":"1196703900","pc_new":6,"screen_name":"猫在飞fly","profile_image_url":"https://tva1.sinaimg.cn/crop.114.19.326.326.50/4754409cjw1f6q2tcgfp2j20f40c2dh0.jpg?KID=imgbed,tva&Expires=1648035252&ssig=Ad4wpprOtu","profile_url":"/u/1196703900","verified":true,"verified_type":0,"domain":"","weihao":"","verified_type_ext":1,"avatar_large":"https://tva1.sinaimg.cn/crop.114.19.326.326.180/4754409cjw1f6q2tcgfp2j20f40c2dh0.jpg?KID=imgbed,tva&Expires=1648035252&ssig=adj9eqrcxH","avatar_hd":"https://tva1.sinaimg.cn/crop.114.19.326.326.1024/4754409cjw1f6q2tcgfp2j20f40c2dh0.jpg?KID=imgbed,tva&Expires=1648035252&ssig=RY6mEvtAhk","follow_me":false,"following":true,"mbrank":7,"mbtype":12,"planet_video":false,"icon_list":[{"type":"vip","data":{"mbrank":7,"mbtype":12}},{"type":"icon","data":{"value":"1","icon_img":"http://n.sinaimg.cn/www/ea31a93f/20220213/feed_icon_BingDwenDwenShueyRhonRhon@2x.png","title":"微博实现一墩一融","url":"https://s.weibo.com/weibo?q=%23为中国冬奥健儿加油%23"}}]},"can_edit":false,"text_raw":"回复@隔日知输赢:能赚一顿半肯德基吧//@隔日知输赢:老师，重银配债配了6手，想听听您对重债看法[鲜花]//@猫在飞fly:回复@_BELAMI_:不知道，不好意思呢~//@_BELAMI_:这个还会往上涨涨吗？下午比上午低了好多","text":"回复<a href=/n/隔日知输赢>@隔日知输赢</a>:能赚一顿半肯德基吧//<a href=/n/隔日知输赢>@隔日知输赢</a>:老师，重银配债配了6手，想听听您对重债看法<img alt=\"[鲜花]\" title=\"[鲜花]\" src=\"https://face.t.sinajs.cn/t4/appstyle/expression/ext/normal/d4/2018new_xianhua_org.png\" />//<a href=/n/猫在飞fly>@猫在飞fly</a>:回复<a href=/n/_BELAMI_>@_BELAMI_</a>:不知道，不好意思呢~//<a href=/n/_BELAMI_>@_BELAMI_</a>:这个还会往上涨涨吗？下午比上午低了好多","source":"微博 weibo.com","favorited":false,"rid":"6_0_200_6665576152895923232_0_0_0","cardid":"star_1230","is_controlled_by_server":"0","pic_ids":[],"geo":null,"pic_num":0,"is_paid":false,"pic_bg_new":"http://vip.storage.weibo.com/feed_cover/star_1230_mobile_new.png?version=2021091501","mblog_vip_type":0,"number_display_strategy":{"apply_scenario_flag":3,"display_text_min_number":1000000,"display_text":"100万+"},"reposts_count":0,"comments_count":15,"attitudes_count":42,"attitudes_status":0,"isLongText":false,"mlevel":0,"content_auth":0,"is_show_bulletin":2,"comment_manage_info":{"comment_permission_type":-1,"approval_comment_type":0,"comment_sort_type":0},"repost_type":4,"share_repost_type":0,"mblogtype":0,"showFeedRepost":false,"showFeedComment":false,"pictureViewerSign":false,"showPictureViewer":false,"rcList":[],"customIcons":[],"retweeted_status":{"visible":{"type":0,"list_id":0},"created_at":"Wed Mar 23 10:10:55 +0800 2022","id":4750124746543412,"idstr":"4750124746543412","mid":"4750124746543412","mblogid":"Ll1YmrseU","user":{"id":1196703900,"idstr":"1196703900","pc_new":6,"screen_name":"猫在飞fly","profile_image_url":"https://tva1.sinaimg.cn/crop.114.19.326.326.50/4754409cjw1f6q2tcgfp2j20f40c2dh0.jpg?KID=imgbed,tva&Expires=1648035252&ssig=Ad4wpprOtu","profile_url":"/u/1196703900","verified":true,"verified_type":0,"domain":"","weihao":"","verified_type_ext":1,"avatar_large":"https://tva1.sinaimg.cn/crop.114.19.326.326.180/4754409cjw1f6q2tcgfp2j20f40c2dh0.jpg?KID=imgbed,tva&Expires=1648035252&ssig=adj9eqrcxH","avatar_hd":"https://tva1.sinaimg.cn/crop.114.19.326.326.1024/4754409cjw1f6q2tcgfp2j20f40c2dh0.jpg?KID=imgbed,tva&Expires=1648035252&ssig=RY6mEvtAhk","follow_me":false,"following":true,"mbrank":7,"mbtype":12,"planet_video":false,"icon_list":[{"type":"vip","data":{"mbrank":7,"mbtype":12}},{"type":"icon","data":{"value":"1","icon_img":"http://n.sinaimg.cn/www/ea31a93f/20220213/feed_icon_BingDwenDwenShueyRhonRhon@2x.png","title":"微博实现一墩一融","url":"https://s.weibo.com/weibo?q=%23为中国冬奥健儿加油%23"}}]},"can_edit":false,"text_raw":"华友和预期的差了些，基本ok ​​​","text":"华友和预期的差了些，基本ok ​​​","textLength":26,"source":"微博 weibo.com","favorited":false,"rid":"6_0_200_6665576152895923232_0_0_0","cardid":"star_1230","pic_ids":["4754409cly1h0jleiy7jlj20cj0b0my8"],"geo":null,"pic_num":1,"pic_infos":{"4754409cly1h0jleiy7jlj20cj0b0my8":{"thumbnail":{"url":"https://wx4.sinaimg.cn/wap180/4754409cly1h0jleiy7jlj20cj0b0my8.jpg","width":180,"height":158,"cut_type":1,"type":null},"bmiddle":{"url":"https://wx4.sinaimg.cn/wap360/4754409cly1h0jleiy7jlj20cj0b0my8.jpg","width":360,"height":316,"cut_type":1,"type":null},"large":{"url":"https://wx4.sinaimg.cn/orj960/4754409cly1h0jleiy7jlj20cj0b0my8.jpg","width":"451","height":"396","cut_type":1,"type":null},"original":{"url":"https://wx4.sinaimg.cn/orj1080/4754409cly1h0jleiy7jlj20cj0b0my8.jpg","width":"451","height":"396","cut_type":1,"type":null},"largest":{"url":"https://wx4.sinaimg.cn/large/4754409cly1h0jleiy7jlj20cj0b0my8.jpg","width":"451","height":"396","cut_type":1,"type":null},"mw2000":{"url":"https://wx4.sinaimg.cn/mw2000/4754409cly1h0jleiy7jlj20cj0b0my8.jpg","width":"451","height":"396","cut_type":1,"type":null},"object_id":"1042018:736d1845d5c38369235ac4766f50f820","pic_id":"4754409cly1h0jleiy7jlj20cj0b0my8","photo_tag":0,"type":"pic","pic_status":0}},"is_paid":false,"mblog_vip_type":0,"number_display_strategy":{"apply_scenario_flag":3,"display_text_min_number":1000000,"display_text":"100万+"},"reposts_count":4,"comments_count":35,"attitudes_count":144,"attitudes_status":0,"isLongText":false,"mlevel":0,"content_auth":0,"is_show_bulletin":2,"comment_manage_info":{"comment_permission_type":-1,"approval_comment_type":0,"comment_sort_type":0},"mblogtype":0,"showFeedRepost":false,"showFeedComment":false,"pictureViewerSign":false,"showPictureViewer":false,"rcList":[],"customIcons":[]}},{"visible":{"type":0,"list_id":0},"created_at":"Wed Mar 23 15:17:02 +0800 2022","id":4750201783322313,"idstr":"4750201783322313","mid":"4750201783322313","mblogid":"Ll3YCdWhH","user":{"id":1402070602,"idstr":"1402070602","pc_new":7,"screen_name":"诺饭短线","profile_image_url":"https://tva1.sinaimg.cn/crop.0.0.180.180.50/5391e64ajw1e8qgp5bmzyj2050050aa8.jpg?KID=imgbed,tva&Expires=1648035252&ssig=NpmsWD8KPj","profile_url":"/u/1402070602","verified":true,"verified_type":0,"domain":"swgreat","weihao":"275745544","verified_type_ext":1,"avatar_large":"https://tva1.sinaimg.cn/crop.0.0.180.180.180/5391e64ajw1e8qgp5bmzyj2050050aa8.jpg?KID=imgbed,tva&Expires=1648035252&ssig=QrIREWj3n0","avatar_hd":"https://tva1.sinaimg.cn/crop.0.0.180.180.1024/5391e64ajw1e8qgp5bmzyj2050050aa8.jpg?KID=imgbed,tva&Expires=1648035252&ssig=SifMXru1Ea","follow_me":false,"following":true,"mbrank":6,"mbtype":11,"planet_video":true,"icon_list":[{"type":"vip","data":{"mbrank":6,"mbtype":11}}]},"can_edit":false,"text_raw":"尾盘跌的我已经有点恐慌了，恐慌并不是因为炸板潮或者中国医药的大跌，这个我早就做过预案，很可能会有一个双中阴的下跌，所以我昨天没有补。恐慌是因为明天万一不反弹怎么办。万一走成准油那种下跌怎么办。所以尾盘其实不该卖的南威和岭南我也卖掉了，已经管不了那么多了。因为在我们的模式里面，总龙 ​​​","text":"尾盘跌的我已经有点恐慌了，恐慌并不是因为炸板潮或者中国医药的大跌，这个我早就做过预案，很可能会有一个双中阴的下跌，所以我昨天没有补。恐慌是因为明天万一不反弹怎么办。万一走成准油那种下跌怎么办。所以尾盘其实不该卖的南威和岭南我也卖掉了，已经管不了那么多了。因为在我们的模式里面，总龙 ​​​ ...<span class=\"expand\">展开</span>","textLength":984,"source":"新版微博 weibo.com","favorited":false,"rid":"7_0_200_6665576152895923232_0_0_0","is_controlled_by_server":"0","pic_ids":[],"geo":null,"pic_num":0,"is_paid":false,"mblog_vip_type":0,"number_display_strategy":{"apply_scenario_flag":3,"display_text_min_number":1000000,"display_text":"100万+"},"reposts_count":2,"comments_count":29,"attitudes_count":40,"attitudes_status":0,"continue_tag":{"title":"全文","pic":"http://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article.png","scheme":"sinaweibo://detail?mblogid=4750201783322313&id=4750201783322313"},"isLongText":true,"mlevel":0,"content_auth":0,"is_show_bulletin":2,"comment_manage_info":{"comment_permission_type":-1,"approval_comment_type":0,"comment_sort_type":0},"share_repost_type":0,"mblogtype":0,"showFeedRepost":false,"showFeedComment":false,"pictureViewerSign":false,"showPictureViewer":false,"rcList":[],"customIcons":[]},{"visible":{"type":0,"list_id":0},"created_at":"Wed Mar 23 15:10:31 +0800 2022","id":4750200147280150,"idstr":"4750200147280150","mid":"4750200147280150","mblogid":"Ll3VYuxTM","user":{"id":3612749480,"idstr":"3612749480","pc_new":6,"screen_name":"山人I","profile_image_url":"https://tvax3.sinaimg.cn/crop.0.0.828.828.50/d7562ea8ly8gnyu8bneyoj20n00n0jt8.jpg?KID=imgbed,tva&Expires=1648035252&ssig=0jrxHNyYN6","profile_url":"/u/3612749480","verified":true,"verified_type":0,"domain":"","weihao":"","verified_type_ext":1,"avatar_large":"https://tvax3.sinaimg.cn/crop.0.0.828.828.180/d7562ea8ly8gnyu8bneyoj20n00n0jt8.jpg?KID=imgbed,tva&Expires=1648035252&ssig=Lq%2BaZDZh9G","avatar_hd":"https://tvax3.sinaimg.cn/crop.0.0.828.828.1024/d7562ea8ly8gnyu8bneyoj20n00n0jt8.jpg?KID=imgbed,tva&Expires=1648035252&ssig=YXtkeuJPA5","follow_me":false,"following":true,"mbrank":7,"mbtype":12,"planet_video":false,"icon_list":[{"type":"vip","data":{"mbrank":7,"mbtype":12}}]},"can_edit":false,"edit_count":1,"text_raw":"今天场外的净值是0，惊不惊喜？\n我快傻了，关着家里人都木了。 ​​​","text":"今天场外的净值是0，惊不惊喜？<br />我快傻了，关着家里人都木了。 ​​​","textLength":58,"source":"iPhone 11","favorited":false,"rid":"8_0_200_6665576152895923232_0_0_0","is_controlled_by_server":"0","pic_ids":["d7562ea8gy1h0ju65not6g209q05qqud"],"geo":null,"pic_num":1,"pic_infos":{"d7562ea8gy1h0ju65not6g209q05qqud":{"thumbnail":{"url":"https://wx3.sinaimg.cn/wap180/d7562ea8gy1h0ju65not6g209q05qqud.gif","width":180,"height":105,"cut_type":1,"type":null},"bmiddle":{"url":"https://wx3.sinaimg.cn/wap360/d7562ea8gy1h0ju65not6g209q05qqud.gif","width":"350","height":"206","cut_type":1,"type":null},"large":{"url":"https://wx3.sinaimg.cn/orj960/d7562ea8gy1h0ju65not6g209q05qqud.gif","width":"350","height":"206","cut_type":1,"type":null},"original":{"url":"https://wx3.sinaimg.cn/orj1080/d7562ea8gy1h0ju65not6g209q05qqud.gif","width":"350","height":"206","cut_type":1,"type":null},"largest":{"url":"https://wx3.sinaimg.cn/large/d7562ea8gy1h0ju65not6g209q05qqud.gif","width":"350","height":"206","cut_type":1,"type":null},"mw2000":{"url":"https://wx3.sinaimg.cn/mw2000/d7562ea8gy1h0ju65not6g209q05qqud.gif","width":"350","height":"206","cut_type":1,"type":null},"object_id":"1042018:bd4ab66addf27730f9a46c896199f02f","pic_id":"d7562ea8gy1h0ju65not6g209q05qqud","photo_tag":0,"type":"gif","video_object_id":"1022:2311284750201085427744","video":"http://g.us.sinaimg.cn/o0/000fCxO4lx07UHHYQZLa010412000gnU0E010.mp4?label=gif_mp4&template=350x206.28.0&ssig=41vZ38vh3E&Expires=1648028052&KID=unistore,video","pic_status":1}},"is_paid":false,"mblog_vip_type":0,"number_display_strategy":{"apply_scenario_flag":3,"display_text_min_number":1000000,"display_text":"100万+"},"reposts_count":0,"comments_count":213,"attitudes_count":638,"attitudes_status":0,"isLongText":false,"mlevel":0,"content_auth":0,"is_show_bulletin":2,"comment_manage_info":{"comment_permission_type":-1,"approval_comment_type":0,"comment_sort_type":0},"share_repost_type":0,"mblogtype":0,"showFeedRepost":false,"showFeedComment":false,"pictureViewerSign":false,"showPictureViewer":false,"rcList":[],"customIcons":[]},{"visible":{"type":0,"list_id":0},"created_at":"Wed Mar 23 15:03:26 +0800 2022","id":4750198359722526,"idstr":"4750198359722526","mid":"4750198359722526","mblogid":"Ll3T5ENgW","user":{"id":5589534463,"idstr":"5589534463","pc_new":0,"screen_name":"创业板专业户1600","profile_image_url":"https://tvax3.sinaimg.cn/crop.128.44.399.399.50/0066h7fFly8gd7bql5uoej30fq0cm75f.jpg?KID=imgbed,tva&Expires=1648035252&ssig=bDVzzaKeYd","profile_url":"/u/5589534463","verified":true,"verified_type":0,"domain":"","weihao":"","verified_type_ext":0,"avatar_large":"https://tvax3.sinaimg.cn/crop.128.44.399.399.180/0066h7fFly8gd7bql5uoej30fq0cm75f.jpg?KID=imgbed,tva&Expires=1648035252&ssig=RNemjAs%2Fck","avatar_hd":"https://tvax3.sinaimg.cn/crop.128.44.399.399.1024/0066h7fFly8gd7bql5uoej30fq0cm75f.jpg?KID=imgbed,tva&Expires=1648035252&ssig=AKCKS%2FA3Tt","follow_me":false,"following":true,"mbrank":2,"mbtype":2,"planet_video":true,"icon_list":[]},"can_edit":false,"text_raw":"这波创业板反弹，核心是超跌的医药cro,去年大跌后，反弹一个半月了，但也反弹到前期缺口压力位置了，后面必须有新方向出来，否则创业板也拉不上2740强压力位 ​​​","text":"这波创业板反弹，核心是超跌的医药cro,去年大跌后，反弹一个半月了，但也反弹到前期缺口压力位置了，后面必须有新方向出来，否则创业板也拉不上2740强压力位 ​​​","textLength":146,"source":"东方财富手机版","favorited":false,"rid":"9_0_200_6665576152895923232_0_0_0","is_controlled_by_server":"0","pic_ids":["0066h7fFgy1h0jtsfs1ujj30u027vtft"],"geo":null,"pic_num":1,"pic_infos":{"0066h7fFgy1h0jtsfs1ujj30u027vtft":{"thumbnail":{"url":"https://wx3.sinaimg.cn/wap180/0066h7fFgy1h0jtsfs1ujj30u027vtft.jpg","width":67,"height":180,"cut_type":1,"type":null},"bmiddle":{"url":"https://wx3.sinaimg.cn/wap360/0066h7fFgy1h0jtsfs1ujj30u027vtft.jpg","width":135,"height":360,"cut_type":1,"type":null},"large":{"url":"https://wx3.sinaimg.cn/orj960/0066h7fFgy1h0jtsfs1ujj30u027vtft.jpg","width":960,"height":2555,"cut_type":1,"type":null},"original":{"url":"https://wx3.sinaimg.cn/orj1080/0066h7fFgy1h0jtsfs1ujj30u027vtft.jpg","width":"1080","height":"2875","cut_type":1,"type":null},"largest":{"url":"https://wx3.sinaimg.cn/large/0066h7fFgy1h0jtsfs1ujj30u027vtft.jpg","width":"1080","height":"2875","cut_type":1,"type":null},"mw2000":{"url":"https://wx3.sinaimg.cn/mw2000/0066h7fFgy1h0jtsfs1ujj30u027vtft.jpg","width":"1080","height":"2875","cut_type":1,"type":null},"object_id":"1042018:5d851aa29c267cf1b0ae792ff3e5dfb0","pic_id":"0066h7fFgy1h0jtsfs1ujj30u027vtft","photo_tag":0,"type":"pic","pic_status":0}},"is_paid":false,"mblog_vip_type":0,"number_display_strategy":{"apply_scenario_flag":3,"display_text_min_number":1000000,"display_text":"100万+"},"reposts_count":1,"comments_count":0,"attitudes_count":28,"attitudes_status":0,"isLongText":false,"mlevel":0,"content_auth":0,"is_show_bulletin":2,"comment_manage_info":{"comment_permission_type":-1,"approval_comment_type":0,"comment_sort_type":0},"share_repost_type":0,"tag_struct":[{"tag_name":"创业板指 2715.95","oid":"1022:230677sz399006","tag_type":2,"tag_hidden":0,"tag_scheme":"https://stock.weibo.cn/page/tag?oid=1022:230677sz399006","url_type_pic":"https://h5.sinaimg.cn/upload/1008/253/2018/12/26/timeline_icon_stock.png","actionlog":{"act_code":2413,"oid":"1022:230677sz399006","uicode":null,"luicode":null,"fid":null,"ext":"|tag_type:stock"},"bd_object_type":"stock","desc":"+1.03%"}],"mblogtype":0,"showFeedRepost":false,"showFeedComment":false,"pictureViewerSign":false,"showPictureViewer":false,"rcList":[],"customIcons":[]}],"total_number":2000,"since_id":4750217209449536,"max_id":4750196611224356,"since_id_str":"4750217209449536","max_id_str":"4750196611224356"}`

var data2 = `
{
"data": {
"list": [
{
"business_type": 1,
"create_time": "2022-01-21 15:35:44",
"credit_rate": "A1",
"customer_create_time": "2019-01-30 15:39:54",
"customer_id": 14,
"customer_type": 1,
"done_time": "2022-01-21 15:35:46",
"id": 39,
"max_amount": 1500000,
"run_type": 1,
"source": 2,
"stage": 5
},
{
"business_type": 1,
"create_time": "2022-01-21 15:32:36",
"credit_rate": "A1",
"customer_create_time": "2019-01-30 15:39:54",
"customer_id": 14,
"customer_type": 1,
"done_time": "2022-01-21 15:32:38",
"id": 38,
"max_amount": 1500000,
"run_type": 2,
"source": 2,
"stage": 5
},
{
"business_type": 1,
"create_time": "2022-01-21 10:09:08",
"credit_rate": "",
"customer_create_time": "2019-01-30 15:39:54",
"customer_id": 14,
"customer_type": 1,
"done_time": "",
"id": 37,
"max_amount": 0,
"run_type": 2,
"source": 2,
"stage": 2
},
{
"business_type": 1,
"create_time": "2022-01-21 10:09:01",
"credit_rate": "",
"customer_create_time": "2018-11-15 07:34:53",
"customer_id": 10,
"customer_type": 0,
"done_time": "",
"id": 36,
"max_amount": 0,
"run_type": 2,
"source": 2,
"stage": 2
},
{
"business_type": 1,
"create_time": "2022-01-21 10:08:53",
"credit_rate": "",
"customer_create_time": "2018-11-26 06:39:29",
"customer_id": 6,
"customer_type": 0,
"done_time": "",
"id": 35,
"max_amount": 0,
"run_type": 2,
"source": 2,
"stage": 2
},
{
"business_type": 1,
"create_time": "2022-01-21 10:08:49",
"credit_rate": "",
"customer_create_time": "2018-11-15 07:34:53",
"customer_id": 4,
"customer_type": 1,
"done_time": "",
"id": 34,
"max_amount": 0,
"run_type": 2,
"source": 2,
"stage": 2
},
{
"business_type": 1,
"create_time": "2022-01-21 10:08:17",
"credit_rate": "",
"customer_create_time": "2019-01-30 15:39:54",
"customer_id": 14,
"customer_type": 1,
"done_time": "",
"id": 33,
"max_amount": 0,
"run_type": 2,
"source": 2,
"stage": 2
},
{
"business_type": 1,
"create_time": "2022-01-21 10:01:35",
"credit_rate": "",
"customer_create_time": "2020-03-29 18:24:24",
"customer_id": 2,
"customer_type": 0,
"done_time": "",
"id": 32,
"max_amount": 0,
"run_type": 2,
"source": 2,
"stage": 1
},
{
"business_type": 1,
"create_time": "2022-01-21 09:58:06",
"credit_rate": "",
"customer_create_time": "2020-03-29 18:24:24",
"customer_id": 2,
"customer_type": 0,
"done_time": "",
"id": 31,
"max_amount": 0,
"run_type": 2,
"source": 2,
"stage": 1
},
{
"business_type": 1,
"create_time": "2022-01-21 09:53:34",
"credit_rate": "",
"customer_create_time": "2020-03-29 18:24:24",
"customer_id": 2,
"customer_type": 0,
"done_time": "",
"id": 30,
"max_amount": 0,
"run_type": 2,
"source": 2,
"stage": 1
},
{
"business_type": 1,
"create_time": "2022-01-21 09:52:40",
"credit_rate": "",
"customer_create_time": "2020-03-29 18:24:24",
"customer_id": 2,
"customer_type": 0,
"done_time": "",
"id": 29,
"max_amount": 0,
"run_type": 2,
"source": 2,
"stage": 1
},
{
"business_type": 1,
"create_time": "2022-01-21 09:46:00",
"credit_rate": "",
"customer_create_time": "2019-01-30 15:39:54",
"customer_id": 14,
"customer_type": 1,
"done_time": "",
"id": 28,
"max_amount": 0,
"run_type": 2,
"source": 2,
"stage": 2
},
{
"business_type": 1,
"create_time": "2022-01-21 04:41:18",
"credit_rate": "",
"customer_create_time": "2019-01-30 15:39:54",
"customer_id": 14,
"customer_type": 1,
"done_time": "",
"id": 27,
"max_amount": 0,
"run_type": 2,
"source": 2,
"stage": 2
},
{
"business_type": 1,
"create_time": "2022-01-21 04:32:13",
"credit_rate": "",
"customer_create_time": "2019-01-30 15:39:54",
"customer_id": 14,
"customer_type": 1,
"done_time": "",
"id": 26,
"max_amount": 0,
"run_type": 2,
"source": 2,
"stage": 2
},
{
"business_type": 1,
"create_time": "2022-01-21 04:29:41",
"credit_rate": "",
"customer_create_time": "2019-01-30 15:39:54",
"customer_id": 14,
"customer_type": 1,
"done_time": "",
"id": 25,
"max_amount": 0,
"run_type": 2,
"source": 2,
"stage": 2
}
],
"page": {
"count": 39,
"current_page": 1,
"page_size": 15
}
},
"message": "ok",
"status": 1
}
`

func JsomM(data map[string]interface{}) []interface{} {
	var array []interface{}

	for _, d := range data {
		b, err2 := json.Marshal(d)
		if err2 != nil {
			fmt.Println("json.Marshal failed:", err2)
			d2 := d.(map[string]interface{})
			JsomM(d2)
			continue
		}
		var ret Article
		err3 := json.Unmarshal([]byte(b), &ret)
		if err3 != nil {
			continue

		}
		if ret.ArticleId == "" {
			continue
		}

		array = append(array, ret)
	}

	return array
}

func TestJson2(t *testing.T) {
	//var v interface{}
	//var v JsonData

	var a Article
	////GetJsonTag(a)
	//
	//
	var strIndex []int

	tM1 := make(map[int]string)
	sM1 := make(map[int]string)
	pM1 := make(map[string]interface{})
	val := reflect.ValueOf(a)
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		sf := typ.Field(i)
		tag := sf.Tag.Get("json")
		if tag == "-" {
			continue
		}
		tM1[int(sf.Offset)] = fmt.Sprint(sf.Type)
		sM1[int(sf.Offset)] = tag
		strIndex = append(strIndex, int(sf.Offset))
	}

	dest := (uintptr)(unsafe.Pointer(&a))
	for _, b := range strIndex {
		teamp := *(*interface{})(unsafe.Pointer(dest + uintptr(b)))
		pM1[sM1[b]] = &teamp

		regexpS := "\"" + sM1[b] + "\"" + `:(\S+?)((,"[a-zA-Z_]+":)|(},))`

		var RegexAmount = regexp.MustCompile(regexpS)
		data += ","
		newS := strings.Replace(strings.TrimSpace(data), " ", "", -1)
		newS = strings.Replace(strings.TrimSpace(newS), "\n", "", -1)
		match := RegexAmount.FindStringSubmatch(newS)
		if len(match) < 2 {
			continue
		}
		value := match[1]
		switch tM1[b] {
		case "int64":
			fmt.Println("int64")
			amount, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return
			}
			*(*int64)(unsafe.Pointer(dest + uintptr(b))) = amount
		case "string":
			fmt.Println("string")
			*(*string)(unsafe.Pointer(dest + uintptr(b))) = strings.Replace(strings.TrimSpace(value), "\"", "", -1)
		}
	}
	fmt.Println("match:", a)
	//var aa Article
	//b := JsonRegexp(data, a,(uintptr)(unsafe.Pointer(&aa)))
	//
	////fmt.Println("match:", 	*(*Article)(unsafe.Pointer(b)))
	//fmt.Println("match:", aa)
	//fmt.Println("match:", b)
}

func TestJson4(t *testing.T) {
	f := func(data []byte) (interface{}, error) {
		var ret *Article
		err3 := json.Unmarshal(data, &ret)
		if err3 != nil || ret.ArticleId == "" || ret.BriefContent == "" {
			return nil, errors.New("空字段")

		}

		return ret, nil
	}
	array := JsonToDataList(data, f)
	for _, v := range array {
		fmt.Println("v:", v.(*Article))
	}
}

func TestJsonWb(t *testing.T) {
	f := func(data []byte) (interface{}, error) {
		var ret *WbArticle
		err3 := json.Unmarshal(data, &ret)
		if err3 != nil {
			return nil, errors.New("空字段")

		}

		return ret, nil
	}
	array := JsonToDataList(webData, f)
	for _, v := range array {
		fmt.Println("v:", v.(*WbArticle))
	}
}

type S2 struct {
	CustomerCreateTime string `json:"customer_create_time"`
	CustomerId         int    `json:"customer_id"`
	CustomerType       int    `json:"customer_type"`
}

func TestJsonS2(t *testing.T) {
	f := func(data []byte) (interface{}, error) {
		var ret *S2
		err3 := json.Unmarshal(data, &ret)
		if err3 != nil || (ret != nil && ret.CustomerId == 0) {
			return nil, errors.New("空字段")

		}

		return ret, nil
	}
	array := JsonToDataList(data2, f)
	for _, v := range array {
		fmt.Println("v:", v.(*S2))
	}
}

func TestFunc(t *testing.T) {
	aa := []int64{1, 2, 3}
	ids := make([]int64, 0)
	fmt.Println(ids, len(ids), cap(ids))
	for _, v := range aa {
		ids = append(ids, v)
	}

	fmt.Println(ids, len(ids), cap(ids))
}
