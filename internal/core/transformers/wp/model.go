package transformers

// {
// 	"id": 249403,
// 	"date": "2024-02-16T11:58:08",
// 	"date_gmt": "2024-02-16T17:58:08",
// 	"guid": {
// 		"rendered": "https://wsform.com/?post_type=knowledgebase&#038;p=249403"
// 	},
// 	"modified": "2024-02-16T12:00:11",
// 	"modified_gmt": "2024-02-16T18:00:11",
// 	"slug": "how-to-reload-the-page-using-an-action",
// 	"status": "publish",
// 	"type": "knowledgebase",
// 	"link": "https://wsform.com/knowledgebase/how-to-reload-the-page-using-an-action/",
// 	"title": {
// 		"rendered": "How to Reload the Page Using an Action"
// 	},
// 	"content": {
// 		"rendered": "<p>Reload or refreshing a page can be useful if, for example, you are using a form to edit a post and the form is on the post itself. Reloading a page is easy to do using some simple JavaScript. The JavaScript is as follows:</p>\n<pre class=\"EnlighterJSRAW\" data-enlighter-language=\"js\">location.reload();</pre>\n<h2>Using the Run JavaScript Action</h2>\n<p>The <a href=\"https://wsform.com/knowledgebase/run-javascript/\">Run JavaScript</a> action can be used to run the reload method. To do this:</p>\n<ol>\n<li>When editing your form, click the <strong>Actions</strong> <i class=\"icon circle actions\"></i> icon at the top of the page. The <strong>Actions</strong> sidebar will open on the side of the page.</li>\n<li>Click the <strong>Add</strong> <i class=\"icon plus-circle\"></i> icon to add a new action.</li>\n<li>Select <strong>JavaScript</strong> from the Action pulldown.</li>\n<li>Is the <strong>JavaScript</strong> setting, enter <code>location.reload();</code>.</li>\n<li>Click <strong>Save &amp; Close</strong>.</li>\n<li><a href=\"https://wsform.com/knowledgebase/publishing-forms/\">Publish</a> your form.</li>\n</ol>\n<p><img loading=\"lazy\" decoding=\"async\" class=\"size-full wp-image-249407 aligncenter\" src=\"https://wsform.com/wp-content/uploads/2024/02/action_run_javascript_reload.jpg\" alt=\"WS Form - Run JavaScript Action - Reload Page\" width=\"400\" height=\"433\" srcset=\"https://wsform.com/wp-content/uploads/2024/02/action_run_javascript_reload.jpg 400w, https://wsform.com/wp-content/uploads/2024/02/action_run_javascript_reload-277x300.jpg 277w, https://wsform.com/wp-content/uploads/2024/02/action_run_javascript_reload-307x332.jpg 307w, https://wsform.com/wp-content/uploads/2024/02/action_run_javascript_reload-163x176.jpg 163w\" sizes=\"(max-width: 400px) 100vw, 400px\" /></p>\n<p>Now when the form is submitted, the page will reload.</p>\n<p>Note that reloading and redirecting a web page may prevent other JavaScript on the page from running, e.g. analytics events.</p>\n<h2>Using Conditional Logic</h2>\n<p>JavaScript can also be run using <a href=\"https://wsform.com/template/conditional-logic/\">conditional logic</a>. For example you could run JavaScript with a <a href=\"https://wsform.com/knowledgebase/custom/\">custom button</a> is clicked. To do this:</p>\n<ol>\n<li>Click the <strong>Conditional Logic</strong> <i class=\"icon circle conditional\"></i> icon at the top of the layout editor. The Conditional Logic sidebar will open.</li>\n<li>Click the <strong>Add</strong> <i class=\"icon plus-circle\"></i> icon to add a new condition.</li>\n<li>When editing <strong>THEN</strong>, choose <strong>Form &gt; Run JavaScript</strong>.</li>\n<li>Enter <code>location.reload();</code>.</li>\n<li>Click <strong>Save &amp; Close</strong>.</li>\n<li><a href=\"https://wsform.com/knowledgebase/publishing-forms/\">Publish</a> your form.</li>\n</ol>\n<p><img loading=\"lazy\" decoding=\"async\" class=\"size-full wp-image-249409 aligncenter\" src=\"https://wsform.com/wp-content/uploads/2024/02/conditional_logic_run_javascript_reload.jpg\" alt=\"WS Form - Conditional Logic - Run JavaScript - Reload Page\" width=\"400\" height=\"564\" srcset=\"https://wsform.com/wp-content/uploads/2024/02/conditional_logic_run_javascript_reload.jpg 400w, https://wsform.com/wp-content/uploads/2024/02/conditional_logic_run_javascript_reload-213x300.jpg 213w, https://wsform.com/wp-content/uploads/2024/02/conditional_logic_run_javascript_reload-235x332.jpg 235w, https://wsform.com/wp-content/uploads/2024/02/conditional_logic_run_javascript_reload-125x176.jpg 125w\" sizes=\"(max-width: 400px) 100vw, 400px\" /></p>\n<p>In the above example the page would be reloaded each time the custom button is clicked. Clicking the button would clear any data entered into the form at that point.</p>\n<h2>Bypassing The Browser Cache</h2>\n<p>Should you encounter an issue where the page content does not reload correctly, it may be due to the browser cache retrieving a cached version of the page. To address this, you can ensure a fresh reload by setting the <code>forceGet</code> parameter to <code>true</code> in the reload method, as demonstrated below:</p>\n<pre class=\"EnlighterJSRAW\" data-enlighter-language=\"js\">location.reload(true);</pre>\n<p>&nbsp;</p>\n",
// 		"protected": false
// 	},
// 	"excerpt": {
// 		"rendered": "<p>Reloading a page is easy to do using some simple JavaScript and the Run JavaScript action or with conditional logic.</p>\n",
// 		"protected": false
// 	},
// 	"featured_media": 0,
// 	"parent": 0,
// 	"menu_order": 0,
// 	"template": "",
// 	"search_keyword": [
// 		17481,
// 		17480
// 	],
// 	"acf": [],
// 	"yoast_head": "<!-- This site is optimized with the Yoast SEO Premium plugin v21.9 (Yoast SEO v21.9.1) - https://yoast.com/wordpress/plugins/seo/ -->\n<title>How to Reload the Page Using an Action - WS Form</title>\n<meta name=\"robots\" content=\"index, follow, max-snippet:-1, max-image-preview:large, max-video-preview:-1\" />\n<link rel=\"canonical\" href=\"https://wsform.com/knowledgebase/how-to-reload-the-page-using-an-action/\" />\n<meta property=\"og:locale\" content=\"en_US\" />\n<meta property=\"og:type\" content=\"article\" />\n<meta property=\"og:title\" content=\"How to Reload the Page Using an Action\" />\n<meta property=\"og:description\" content=\"Reloading a page is easy to do using some simple JavaScript and the Run JavaScript action or with conditional logic.\" />\n<meta property=\"og:url\" content=\"https://wsform.com/knowledgebase/how-to-reload-the-page-using-an-action/\" />\n<meta property=\"og:site_name\" content=\"WS Form\" />\n<meta property=\"article:publisher\" content=\"https://www.facebook.com/groups/wsform\" />\n<meta property=\"article:modified_time\" content=\"2024-02-16T18:00:11+00:00\" />\n<meta property=\"og:image\" content=\"https://wsform.com/wp-content/uploads/2024/02/action_run_javascript_reload.jpg\" />\n<meta name=\"twitter:card\" content=\"summary_large_image\" />\n<meta name=\"twitter:site\" content=\"@ws_form\" />\n<meta name=\"twitter:label1\" content=\"Est. reading time\" />\n\t<meta name=\"twitter:data1\" content=\"2 minutes\" />\n<script type=\"application/ld+json\" class=\"yoast-schema-graph\">{\"@context\":\"https://schema.org\",\"@graph\":[{\"@type\":\"WebPage\",\"@id\":\"https://wsform.com/knowledgebase/how-to-reload-the-page-using-an-action/\",\"url\":\"https://wsform.com/knowledgebase/how-to-reload-the-page-using-an-action/\",\"name\":\"How to Reload the Page Using an Action - WS Form\",\"isPartOf\":{\"@id\":\"https://wsform.com/#website\"},\"datePublished\":\"2024-02-16T17:58:08+00:00\",\"dateModified\":\"2024-02-16T18:00:11+00:00\",\"breadcrumb\":{\"@id\":\"https://wsform.com/knowledgebase/how-to-reload-the-page-using-an-action/#breadcrumb\"},\"inLanguage\":\"en-US\",\"potentialAction\":[{\"@type\":\"ReadAction\",\"target\":[\"https://wsform.com/knowledgebase/how-to-reload-the-page-using-an-action/\"]}]},{\"@type\":\"BreadcrumbList\",\"@id\":\"https://wsform.com/knowledgebase/how-to-reload-the-page-using-an-action/#breadcrumb\",\"itemListElement\":[{\"@type\":\"ListItem\",\"position\":1,\"name\":\"Home\",\"item\":\"https://wsform.com/\"},{\"@type\":\"ListItem\",\"position\":2,\"name\":\"Knowledge Base\",\"item\":\"https://wsform.com/knowledgebase/\"},{\"@type\":\"ListItem\",\"position\":3,\"name\":\"How to Reload the Page Using an Action\"}]},{\"@type\":\"WebSite\",\"@id\":\"https://wsform.com/#website\",\"url\":\"https://wsform.com/\",\"name\":\"WS Form\",\"description\":\"Build Better WordPress Forms\",\"publisher\":{\"@id\":\"https://wsform.com/#organization\"},\"potentialAction\":[{\"@type\":\"SearchAction\",\"target\":{\"@type\":\"EntryPoint\",\"urlTemplate\":\"https://wsform.com/?s={search_term_string}\"},\"query-input\":\"required name=search_term_string\"}],\"inLanguage\":\"en-US\"},{\"@type\":\"Organization\",\"@id\":\"https://wsform.com/#organization\",\"name\":\"WS Form\",\"url\":\"https://wsform.com/\",\"logo\":{\"@type\":\"ImageObject\",\"inLanguage\":\"en-US\",\"@id\":\"https://wsform.com/#/schema/logo/image/\",\"url\":\"https://wsform.com/wp-content/uploads/2023/08/logo_v1_400.png\",\"contentUrl\":\"https://wsform.com/wp-content/uploads/2023/08/logo_v1_400.png\",\"width\":400,\"height\":114,\"caption\":\"WS Form\"},\"image\":{\"@id\":\"https://wsform.com/#/schema/logo/image/\"},\"sameAs\":[\"https://www.facebook.com/groups/wsform\",\"https://twitter.com/ws_form\"]}]}</script>\n<!-- / Yoast SEO Premium plugin. -->",
// 	"yoast_head_json": {
// 		"title": "How to Reload the Page Using an Action - WS Form",
// 		"robots": {
// 			"index": "index",
// 			"follow": "follow",
// 			"max-snippet": "max-snippet:-1",
// 			"max-image-preview": "max-image-preview:large",
// 			"max-video-preview": "max-video-preview:-1"
// 		},
// 		"canonical": "https://wsform.com/knowledgebase/how-to-reload-the-page-using-an-action/",
// 		"og_locale": "en_US",
// 		"og_type": "article",
// 		"og_title": "How to Reload the Page Using an Action",
// 		"og_description": "Reloading a page is easy to do using some simple JavaScript and the Run JavaScript action or with conditional logic.",
// 		"og_url": "https://wsform.com/knowledgebase/how-to-reload-the-page-using-an-action/",
// 		"og_site_name": "WS Form",
// 		"article_publisher": "https://www.facebook.com/groups/wsform",
// 		"article_modified_time": "2024-02-16T18:00:11+00:00",
// 		"og_image": [
// 			{
// 				"url": "https://wsform.com/wp-content/uploads/2024/02/action_run_javascript_reload.jpg"
// 			}
// 		],
// 		"twitter_card": "summary_large_image",
// 		"twitter_site": "@ws_form",
// 		"twitter_misc": {
// 			"Est. reading time": "2 minutes"
// 		},
// 		"schema": {
// 			"@context": "https://schema.org",
// 			"@graph": [
// 				{
// 					"@type": "WebPage",
// 					"@id": "https://wsform.com/knowledgebase/how-to-reload-the-page-using-an-action/",
// 					"url": "https://wsform.com/knowledgebase/how-to-reload-the-page-using-an-action/",
// 					"name": "How to Reload the Page Using an Action - WS Form",
// 					"isPartOf": {
// 						"@id": "https://wsform.com/#website"
// 					},
// 					"datePublished": "2024-02-16T17:58:08+00:00",
// 					"dateModified": "2024-02-16T18:00:11+00:00",
// 					"breadcrumb": {
// 						"@id": "https://wsform.com/knowledgebase/how-to-reload-the-page-using-an-action/#breadcrumb"
// 					},
// 					"inLanguage": "en-US",
// 					"potentialAction": [
// 						{
// 							"@type": "ReadAction",
// 							"target": [
// 								"https://wsform.com/knowledgebase/how-to-reload-the-page-using-an-action/"
// 							]
// 						}
// 					]
// 				},
// 				{
// 					"@type": "BreadcrumbList",
// 					"@id": "https://wsform.com/knowledgebase/how-to-reload-the-page-using-an-action/#breadcrumb",
// 					"itemListElement": [
// 						{
// 							"@type": "ListItem",
// 							"position": 1,
// 							"name": "Home",
// 							"item": "https://wsform.com/"
// 						},
// 						{
// 							"@type": "ListItem",
// 							"position": 2,
// 							"name": "Knowledge Base",
// 							"item": "https://wsform.com/knowledgebase/"
// 						},
// 						{
// 							"@type": "ListItem",
// 							"position": 3,
// 							"name": "How to Reload the Page Using an Action"
// 						}
// 					]
// 				},
// 				{
// 					"@type": "WebSite",
// 					"@id": "https://wsform.com/#website",
// 					"url": "https://wsform.com/",
// 					"name": "WS Form",
// 					"description": "Build Better WordPress Forms",
// 					"publisher": {
// 						"@id": "https://wsform.com/#organization"
// 					},
// 					"potentialAction": [
// 						{
// 							"@type": "SearchAction",
// 							"target": {
// 								"@type": "EntryPoint",
// 								"urlTemplate": "https://wsform.com/?s={search_term_string}"
// 							},
// 							"query-input": "required name=search_term_string"
// 						}
// 					],
// 					"inLanguage": "en-US"
// 				},
// 				{
// 					"@type": "Organization",
// 					"@id": "https://wsform.com/#organization",
// 					"name": "WS Form",
// 					"url": "https://wsform.com/",
// 					"logo": {
// 						"@type": "ImageObject",
// 						"inLanguage": "en-US",
// 						"@id": "https://wsform.com/#/schema/logo/image/",
// 						"url": "https://wsform.com/wp-content/uploads/2023/08/logo_v1_400.png",
// 						"contentUrl": "https://wsform.com/wp-content/uploads/2023/08/logo_v1_400.png",
// 						"width": 400,
// 						"height": 114,
// 						"caption": "WS Form"
// 					},
// 					"image": {
// 						"@id": "https://wsform.com/#/schema/logo/image/"
// 					},
// 					"sameAs": [
// 						"https://www.facebook.com/groups/wsform",
// 						"https://twitter.com/ws_form"
// 					]
// 				}
// 			]
// 		}
// 	},
// 	"_links": {
// 		"self": [
// 			{
// 				"href": "https://wsform.com/wp-json/wp/v2/knowledgebase/249403"
// 			}
// 		],
// 		"collection": [
// 			{
// 				"href": "https://wsform.com/wp-json/wp/v2/knowledgebase"
// 			}
// 		],
// 		"about": [
// 			{
// 				"href": "https://wsform.com/wp-json/wp/v2/types/knowledgebase"
// 			}
// 		],
// 		"version-history": [
// 			{
// 				"count": 5,
// 				"href": "https://wsform.com/wp-json/wp/v2/knowledgebase/249403/revisions"
// 			}
// 		],
// 		"predecessor-version": [
// 			{
// 				"id": 249411,
// 				"href": "https://wsform.com/wp-json/wp/v2/knowledgebase/249403/revisions/249411"
// 			}
// 		],
// 		"wp:attachment": [
// 			{
// 				"href": "https://wsform.com/wp-json/wp/v2/media?parent=249403"
// 			}
// 		],
// 		"wp:term": [
// 			{
// 				"taxonomy": "search_keyword",
// 				"embeddable": true,
// 				"href": "https://wsform.com/wp-json/wp/v2/search_keyword?post=249403"
// 			}
// 		],
// 		"curies": [
// 			{
// 				"name": "wp",
// 				"href": "https://api.w.org/{rel}",
// 				"templated": true
// 			}
// 		]
// 	}
// },
