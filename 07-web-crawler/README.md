# Web Crawler

We follow webpages and references in the web pages, for 2 levels of recursion.


## Sequential version

This takes about 35 s.
```
➜  07-web-crawler git:(main) ✗ go run 01-sequential/main.go
Count: 1, Depth: 2, Found: http://andcloud.io
Count: 2, Depth: 1, Found: https://medium.com/andcloudio
Count: 3, Depth: 0, Found: https://medium.com/
Count: 4, Depth: 0, Found: https://rsci.app.link/?%24canonical_url=https%3A%2F%2Fmedium.com/andcloudio%3F~feature=LoMobileNavBar&~channel=ShowCollectionHome&~stage=m2
Count: 5, Depth: 0, Found: https://medium.com/m/signin?redirect=https%3A%2F%2Fmedium.com%2Fandcloudio&source=--------------------------nav_reg&operation=login
Count: 6, Depth: 0, Found: https://medium.com/m/signin?redirect=https%3A%2F%2Fmedium.com%2Fandcloudio&source=--------------------------nav_reg&operation=register
Count: 7, Depth: 0, Found: https://twitter.com/andcloudio
Count: 8, Depth: 0, Found: https://medium.com/andcloudio/sase-future-of-enterprise-network-and-network-security-c8dbc095a3f7?source=collection_home---5------0-----------------------
Count: 9, Depth: 0, Found: https://medium.com/@dgunjetti
Count: 10, Depth: 0, Found: https://medium.com/andcloudio/get-visibility-and-control-into-your-saas-apps-usage-c34b7af8688f?source=collection_home---5------1-----------------------
Count: 11, Depth: 0, Found: https://medium.com/andcloudio/visibility-and-policy-monitoring-with-forseti-b223183c763a?source=collection_home---5------2-----------------------
Count: 12, Depth: 0, Found: https://medium.com/andcloudio/edge-security-with-cloud-armor-24de3c6908e6?source=collection_home---5------3-----------------------
Count: 13, Depth: 0, Found: https://medium.com/andcloudio/secure-access-to-web-apps-with-identity-aware-proxy-14b858d9c068?source=collection_home---5------4-----------------------
Count: 14, Depth: 0, Found: https://medium.com/andcloudio/remote-access-with-beyondcorp-f3bedd1432f2?source=collection_home---5------5-----------------------
Count: 15, Depth: 0, Found: https://medium.com/andcloudio/ci-cd-with-cloud-build-adc254b4750f?source=collection_home---5------6-----------------------
Count: 16, Depth: 0, Found: https://medium.com/andcloudio/increase-developer-productivity-with-skaffold-1604a689dc3d?source=collection_home---5------7-----------------------
Count: 17, Depth: 0, Found: https://medium.com/andcloudio/setting-up-development-environment-on-google-cloud-dd91b619cc80?source=collection_home---5------8-----------------------
Count: 18, Depth: 0, Found: https://medium.com/andcloudio/cost-management-with-gcp-part-1-d27fea4d3241?source=collection_home---5------9-----------------------
Count: 19, Depth: 0, Found: https://medium.com/andcloudio/about
Count: 20, Depth: 0, Found: https://medium.com/andcloudio/latest
Count: 21, Depth: 0, Found: https://medium.com/andcloudio/archive
Count: 22, Depth: 0, Found: https://medium.com/about
Count: 23, Depth: 0, Found: https://policy.medium.com/medium-terms-of-service-9db0094a1e0f
Count: 24, Depth: 0, Found: https://policy.medium.com/medium-privacy-policy-f03bf92035c9
Count: 25, Depth: 1, Found: https://medium.com/andcloudio/get-visibility-and-control-into-your-saas-apps-usage-c34b7af8688f
Count: 26, Depth: 0, Found: https://medium.com/?source=post_page-----c34b7af8688f--------------------------------
Count: 27, Depth: 0, Found: https://www.udemy.com/course/concurrency-in-go-golang/?referralCode=5AE5A041D5793C048954
Count: 28, Depth: 0, Found: https://medium.com/andcloudio/tagged/cloud-computing
Count: 29, Depth: 0, Found: https://medium.com/andcloudio/tagged/security
Count: 30, Depth: 0, Found: https://medium.com/andcloudio/tagged/saas
Count: 31, Depth: 0, Found: https://hasawatermark.medium.com/?source=post_internal_links---------0----------------------------
Count: 32, Depth: 0, Found: https://david-artykov.medium.com/?source=post_internal_links---------3----------------------------
Count: 33, Depth: 0, Found: https://amal-hasni.medium.com/?source=post_internal_links---------4----------------------------
Count: 34, Depth: 0, Found: https://billatnapier.medium.com/?source=post_internal_links---------7----------------------------
Count: 35, Depth: 0, Found: https://medium.com/about?autoplay=1&source=post_page-----c34b7af8688f--------------------------------
Count: 36, Depth: 0, Found: https://medium.com/topics?source=post_page-----c34b7af8688f--------------------------------
Count: 37, Depth: 0, Found: https://about.medium.com/creators/?source=post_page-----c34b7af8688f--------------------------------
Count: 38, Depth: 0, Found: https://policy.medium.com/medium-terms-of-service-9db0094a1e0f?source=post_page-----c34b7af8688f--------------------------------
Count: 39, Depth: 0, Found: https://itunes.apple.com/app/medium-everyones-stories/id828256236?pt=698524&mt=8&ct=post_page&source=post_page-----c34b7af8688f--------------------------------
Count: 40, Depth: 0, Found: https://play.google.com/store/apps/details?id=com.medium.reader&source=post_page-----c34b7af8688f--------------------------------
Count: 41, Depth: 1, Found: https://medium.com/andcloudio/edge-security-with-cloud-armor-24de3c6908e6
Count: 42, Depth: 0, Found: https://medium.com/?source=post_page-----24de3c6908e6--------------------------------
Count: 43, Depth: 0, Found: https://medium.com/andcloudio/tagged/web-development
Count: 44, Depth: 0, Found: https://medium.com/andcloudio/tagged/google-cloud-platform
Count: 45, Depth: 0, Found: https://insafali17.medium.com/?source=post_internal_links---------0----------------------------
Count: 46, Depth: 0, Found: https://paggyru.medium.com/database-security-how-to-use-encryption-to-protect-mongodb-data-fad40711919f?source=post_internal_links---------1----------------------------
Count: 47, Depth: 0, Found: https://paggyru.medium.com/?source=post_internal_links---------1----------------------------
Count: 48, Depth: 0, Found: https://ruurtjan.medium.com/dns-propagation-does-not-exist-faa74ec6a53f?source=post_internal_links---------2----------------------------
Count: 49, Depth: 0, Found: https://ruurtjan.medium.com/?source=post_internal_links---------2----------------------------
Count: 50, Depth: 0, Found: https://ismaelfernandezjr.medium.com/?source=post_internal_links---------3----------------------------
Count: 51, Depth: 0, Found: https://javinpaul.medium.com/?source=post_internal_links---------4----------------------------
Count: 52, Depth: 0, Found: https://wordsbyjohn.medium.com/?source=post_internal_links---------7----------------------------
Count: 53, Depth: 0, Found: https://medium.com/about?autoplay=1&source=post_page-----24de3c6908e6--------------------------------
Count: 54, Depth: 0, Found: https://medium.com/topics?source=post_page-----24de3c6908e6--------------------------------
Count: 55, Depth: 0, Found: https://about.medium.com/creators/?source=post_page-----24de3c6908e6--------------------------------
Count: 56, Depth: 0, Found: https://policy.medium.com/medium-terms-of-service-9db0094a1e0f?source=post_page-----24de3c6908e6--------------------------------
Count: 57, Depth: 0, Found: https://itunes.apple.com/app/medium-everyones-stories/id828256236?pt=698524&mt=8&ct=post_page&source=post_page-----24de3c6908e6--------------------------------
Count: 58, Depth: 0, Found: https://play.google.com/store/apps/details?id=com.medium.reader&source=post_page-----24de3c6908e6--------------------------------
Count: 59, Depth: 1, Found: https://medium.com/andcloudio/remote-access-with-beyondcorp-f3bedd1432f2
Count: 60, Depth: 0, Found: https://medium.com/?source=post_page-----f3bedd1432f2--------------------------------
Count: 61, Depth: 0, Found: https://medium.com/andcloudio/tagged/remote-working
Count: 62, Depth: 0, Found: https://medium.com/andcloudio/tagged/cloud
Count: 63, Depth: 0, Found: https://arslanmalik.medium.com/vulnerability-threat-control-paradigm-and-cia-triads-computer-security-637fe50e8370?source=post_internal_links---------0----------------------------
Count: 64, Depth: 0, Found: https://arslanmalik.medium.com/?source=post_internal_links---------0----------------------------
Count: 65, Depth: 0, Found: https://digitallyvicarious.medium.com/?source=post_internal_links---------1----------------------------
Count: 66, Depth: 0, Found: https://billatnapier.medium.com/?source=post_internal_links---------4----------------------------
Count: 67, Depth: 0, Found: https://philsiarri.medium.com/?source=post_internal_links---------5----------------------------
Count: 68, Depth: 0, Found: https://thenaterhood.medium.com/search-neutrality-3e8ebe588583?source=post_internal_links---------6----------------------------
Count: 69, Depth: 0, Found: https://thenaterhood.medium.com/?source=post_internal_links---------6----------------------------
Count: 70, Depth: 0, Found: https://medium.com/about?autoplay=1&source=post_page-----f3bedd1432f2--------------------------------
Count: 71, Depth: 0, Found: https://medium.com/topics?source=post_page-----f3bedd1432f2--------------------------------
Count: 72, Depth: 0, Found: https://about.medium.com/creators/?source=post_page-----f3bedd1432f2--------------------------------
Count: 73, Depth: 0, Found: https://policy.medium.com/medium-terms-of-service-9db0094a1e0f?source=post_page-----f3bedd1432f2--------------------------------
Count: 74, Depth: 0, Found: https://itunes.apple.com/app/medium-everyones-stories/id828256236?pt=698524&mt=8&ct=post_page&source=post_page-----f3bedd1432f2--------------------------------
Count: 75, Depth: 0, Found: https://play.google.com/store/apps/details?id=com.medium.reader&source=post_page-----f3bedd1432f2--------------------------------
Count: 76, Depth: 1, Found: https://twitter.com/@andcloudio
Elapsed Time: 35.335742172s
```


## Concurrent version
Using go-routines, we don't have to wait sequentially on each webpage to be read, and we obtain a quite remarkable speedup. This version takes just 2s.
```
Count: 1, Depth: 2, Found: http://andcloud.io
Count: 2, Depth: 1, Found: https://medium.com/andcloudio/edge-security-with-cloud-armor-24de3c6908e6
Count: 3, Depth: 1, Found: https://medium.com/andcloudio/remote-access-with-beyondcorp-f3bedd1432f2
Count: 4, Depth: 1, Found: https://medium.com/andcloudio/get-visibility-and-control-into-your-saas-apps-usage-c34b7af8688f
Count: 5, Depth: 0, Found: https://medium.com/about?autoplay=1&source=post_page-----c34b7af8688f--------------------------------
Count: 6, Depth: 0, Found: https://medium.com/about?autoplay=1&source=post_page-----24de3c6908e6--------------------------------
Count: 7, Depth: 1, Found: https://twitter.com/@andcloudio
Count: 8, Depth: 1, Found: https://medium.com/andcloudio
Count: 9, Depth: 0, Found: https://medium.com/about?autoplay=1&source=post_page-----f3bedd1432f2--------------------------------
Count: 10, Depth: 0, Found: https://thenaterhood.medium.com/search-neutrality-3e8ebe588583?source=post_internal_links---------6----------------------------
Count: 11, Depth: 0, Found: https://ruurtjan.medium.com/dns-propagation-does-not-exist-faa74ec6a53f?source=post_internal_links---------2----------------------------
Count: 12, Depth: 0, Found: https://policy.medium.com/medium-terms-of-service-9db0094a1e0f?source=post_page-----24de3c6908e6--------------------------------
Count: 13, Depth: 0, Found: https://policy.medium.com/medium-terms-of-service-9db0094a1e0f?source=post_page-----f3bedd1432f2--------------------------------
Count: 14, Depth: 0, Found: https://policy.medium.com/medium-terms-of-service-9db0094a1e0f?source=post_page-----c34b7af8688f--------------------------------
Count: 15, Depth: 0, Found: https://policy.medium.com/medium-privacy-policy-f03bf92035c9
Count: 16, Depth: 0, Found: https://policy.medium.com/medium-terms-of-service-9db0094a1e0f
Count: 17, Depth: 0, Found: https://play.google.com/store/apps/details?id=com.medium.reader&source=post_page-----c34b7af8688f--------------------------------
Count: 18, Depth: 0, Found: https://medium.com/andcloudio/visibility-and-policy-monitoring-with-forseti-b223183c763a?source=collection_home---5------2-----------------------
Count: 19, Depth: 0, Found: https://medium.com/andcloudio/get-visibility-and-control-into-your-saas-apps-usage-c34b7af8688f?source=collection_home---5------1-----------------------
Count: 20, Depth: 0, Found: https://medium.com/?source=post_page-----24de3c6908e6--------------------------------
Count: 21, Depth: 0, Found: https://play.google.com/store/apps/details?id=com.medium.reader&source=post_page-----24de3c6908e6--------------------------------
Count: 22, Depth: 0, Found: https://medium.com/andcloudio/remote-access-with-beyondcorp-f3bedd1432f2?source=collection_home---5------5-----------------------
Count: 23, Depth: 0, Found: https://play.google.com/store/apps/details?id=com.medium.reader&source=post_page-----f3bedd1432f2--------------------------------
Count: 24, Depth: 0, Found: https://medium.com/andcloudio/secure-access-to-web-apps-with-identity-aware-proxy-14b858d9c068?source=collection_home---5------4-----------------------
Count: 25, Depth: 0, Found: https://medium.com/andcloudio/sase-future-of-enterprise-network-and-network-security-c8dbc095a3f7?source=collection_home---5------0-----------------------
Count: 26, Depth: 0, Found: https://medium.com/andcloudio/tagged/web-development
Count: 27, Depth: 0, Found: https://itunes.apple.com/app/medium-everyones-stories/id828256236?pt=698524&mt=8&ct=post_page&source=post_page-----f3bedd1432f2--------------------------------
Count: 28, Depth: 0, Found: https://twitter.com/andcloudio
Count: 29, Depth: 0, Found: https://rsci.app.link/?%24canonical_url=https%3A%2F%2Fmedium.com/andcloudio%3F~feature=LoMobileNavBar&~channel=ShowCollectionHome&~stage=m2
Count: 30, Depth: 0, Found: https://medium.com/andcloudio/tagged/saas
Count: 31, Depth: 0, Found: https://medium.com/andcloudio/tagged/cloud
Count: 32, Depth: 0, Found: https://medium.com/andcloudio/ci-cd-with-cloud-build-adc254b4750f?source=collection_home---5------6-----------------------
Count: 33, Depth: 0, Found: https://medium.com/andcloudio/tagged/remote-working
Count: 34, Depth: 0, Found: https://medium.com/?source=post_page-----f3bedd1432f2--------------------------------
Count: 35, Depth: 0, Found: https://itunes.apple.com/app/medium-everyones-stories/id828256236?pt=698524&mt=8&ct=post_page&source=post_page-----24de3c6908e6--------------------------------
Count: 36, Depth: 0, Found: https://medium.com/andcloudio/increase-developer-productivity-with-skaffold-1604a689dc3d?source=collection_home---5------7-----------------------
Count: 37, Depth: 0, Found: https://medium.com/andcloudio/edge-security-with-cloud-armor-24de3c6908e6?source=collection_home---5------3-----------------------
Count: 38, Depth: 0, Found: https://medium.com/andcloudio/cost-management-with-gcp-part-1-d27fea4d3241?source=collection_home---5------9-----------------------
Count: 39, Depth: 0, Found: https://medium.com/about
Count: 40, Depth: 0, Found: https://medium.com/andcloudio/setting-up-development-environment-on-google-cloud-dd91b619cc80?source=collection_home---5------8-----------------------
Count: 41, Depth: 0, Found: https://medium.com/andcloudio/tagged/cloud-computing
Count: 42, Depth: 0, Found: https://medium.com/andcloudio/tagged/security
Count: 43, Depth: 0, Found: https://medium.com/m/signin?redirect=https%3A%2F%2Fmedium.com%2Fandcloudio&source=--------------------------nav_reg&operation=register
Count: 44, Depth: 0, Found: https://medium.com/topics?source=post_page-----24de3c6908e6--------------------------------
Count: 45, Depth: 0, Found: https://medium.com/topics?source=post_page-----f3bedd1432f2--------------------------------
Count: 46, Depth: 0, Found: https://about.medium.com/creators/?source=post_page-----24de3c6908e6--------------------------------
Count: 47, Depth: 0, Found: https://medium.com/andcloudio/about
Count: 48, Depth: 0, Found: https://about.medium.com/creators/?source=post_page-----f3bedd1432f2--------------------------------
Count: 49, Depth: 0, Found: https://medium.com/andcloudio/tagged/google-cloud-platform
Count: 50, Depth: 0, Found: https://about.medium.com/creators/?source=post_page-----c34b7af8688f--------------------------------
Count: 51, Depth: 0, Found: https://medium.com/topics?source=post_page-----c34b7af8688f--------------------------------
Count: 52, Depth: 0, Found: https://medium.com/
Count: 53, Depth: 0, Found: https://medium.com/andcloudio/latest
Count: 54, Depth: 0, Found: https://medium.com/andcloudio/archive
Count: 55, Depth: 0, Found: https://digitallyvicarious.medium.com/?source=post_internal_links---------1----------------------------
Count: 56, Depth: 0, Found: https://arslanmalik.medium.com/vulnerability-threat-control-paradigm-and-cia-triads-computer-security-637fe50e8370?source=post_internal_links---------0----------------------------
Count: 57, Depth: 0, Found: https://medium.com/?source=post_page-----c34b7af8688f--------------------------------
Count: 58, Depth: 0, Found: https://billatnapier.medium.com/?source=post_internal_links---------7----------------------------
Count: 59, Depth: 0, Found: https://medium.com/m/signin?redirect=https%3A%2F%2Fmedium.com%2Fandcloudio&source=--------------------------nav_reg&operation=login
Count: 60, Depth: 0, Found: https://paggyru.medium.com/database-security-how-to-use-encryption-to-protect-mongodb-data-fad40711919f?source=post_internal_links---------1----------------------------
Count: 61, Depth: 0, Found: https://arslanmalik.medium.com/?source=post_internal_links---------0----------------------------
Count: 62, Depth: 0, Found: https://wordsbyjohn.medium.com/?source=post_internal_links---------7----------------------------
Count: 63, Depth: 0, Found: https://philsiarri.medium.com/?source=post_internal_links---------5----------------------------
Count: 64, Depth: 0, Found: https://insafali17.medium.com/?source=post_internal_links---------0----------------------------
Count: 65, Depth: 0, Found: https://david-artykov.medium.com/?source=post_internal_links---------3----------------------------
Count: 66, Depth: 0, Found: https://billatnapier.medium.com/?source=post_internal_links---------4----------------------------
Count: 67, Depth: 0, Found: https://hasawatermark.medium.com/?source=post_internal_links---------0----------------------------
Count: 68, Depth: 0, Found: https://paggyru.medium.com/?source=post_internal_links---------1----------------------------
Count: 69, Depth: 0, Found: https://amal-hasni.medium.com/?source=post_internal_links---------4----------------------------
Count: 70, Depth: 0, Found: https://thenaterhood.medium.com/?source=post_internal_links---------6----------------------------
Count: 71, Depth: 0, Found: https://javinpaul.medium.com/?source=post_internal_links---------4----------------------------
Count: 72, Depth: 0, Found: https://itunes.apple.com/app/medium-everyones-stories/id828256236?pt=698524&mt=8&ct=post_page&source=post_page-----c34b7af8688f--------------------------------
Count: 73, Depth: 0, Found: https://ruurtjan.medium.com/?source=post_internal_links---------2----------------------------
Count: 74, Depth: 0, Found: https://ismaelfernandezjr.medium.com/?source=post_internal_links---------3----------------------------
Count: 75, Depth: 0, Found: https://medium.com/@dgunjetti
Count: 76, Depth: 0, Found: https://www.udemy.com/course/concurrency-in-go-golang/?referralCode=5AE5A041D5793C048954
Elapsed Time: 2.032099684s
```