
---
label: til
layout: default
title: k8s networking
---
clusterip là chuyện map 1 cái svc ip+port ra nhiều pod ip+port
Cái đó tầng trên
Câu hỏi make sense hơn nếu phrase theo kiểu pod network map ra network ở dưới thế nào
Cái này tuỳ network ở dưới thế nào
Chạy trên L4 thì phải overlay, vd gói L3/L2 trong UDP
Làm dở thì có userspace process listen trên cái port UDP
Gói/mở ở 2 đầu
Làm tốt hơn thì để trong kernel, data path không đi qua process ngoài (list ai listen cái port không ra)
Chạy trên L3 thì có thể dùng IP routing
Xuống tới đây thì khái niệm “mở port” cũng không make sense nữa
Gói IP tới node, nó mở ra nó xem tới pod IP+port nào nó đưa tới thôi
Còn giữa các node với nhau thì phải cấu hình routing (local route trên node, hoặc trên gateway) (edited) 
Cho pod IP
Ít máy physical thì cho static luôn
Hay thêm bớt thì phải thêm cơ chế trao đổi routing state (bgp hay gì đó)
Hoặc là chơi map IP -> pod IP theo kiểu deterministic
e.g. 172.16.a.b -> 10.a.b.0/8
L3 không biết AWS/GCP hỗ trợ thế nào
Nghe nói GCP có
Còn có kiểu chạy trên L2, ipvlan nhiều IP chung trên 1 cái mac
Cái này thì chắc cloud không có rồi
Còn kiểu nhà giàu cho mỗi pod 1 cái eni hoặc physical card luôn

