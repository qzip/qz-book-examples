package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
)

func main() {
	jz := make(map[string]interface{})
	if err := json.Unmarshal([]byte(data), &jz); err != nil {
		log.Fatal(err)
		return
	}
	tx, err := template.New("wellness").Parse(tmplr)
	if err != nil {
		log.Fatal(err)
		return
	}
	var b bytes.Buffer
	if err = tx.Execute(&b, jz); err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(b.String())
}

const data = `
{
  "lastReviewEntry": "2020-01-30",
  "lastJournalEntry":"2020-02-03",
  "lastHealthEntry":"2020-02-04",
  "quesAnsDate":"2020-01-02",
  "signupDate":"2020-01-01",
  "wellYouPoint": 100,
  "wellnessScore":"20",
  "department":"Pschology",
  "university":"Univ. of Delhi",
  "location":"Delhi",
  "dob":"1980-12-18",
  "gender":"male",
  "email":"poc@example.com",
  "firstName":"Moni",
  "lastName":"Baba"
  

}
`

const tmplr = `
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>
        
    Your Profile

    </title>
<!-- <link href="https://fonts.googleapis.com/css?family=Roboto&display=swap" rel="stylesheet"> -->
 <link href="https://fonts.googleapis.com/css?family=Roboto&display=swap" rel="stylesheet">
    <style>
      body {
        font-family: 'Roboto', serif;
        font-size: 48px;
      }
       @font-face {
            src: "https://fonts.googleapis.com/css?family=Roboto";
            font-family: 'Roboto', serif;
            font-style: normal;
            font-weight: normal;
        }

        html {
            font-family: sans-serif;
            -ms-text-size-adjust: 100%;
            -webkit-text-size-adjust: 100%
        }

        body {
            margin: 0;
            background-color: #e9ebee;
             font-family: 'Roboto', serif;
            line-height: 1.5;
        }

        #top-bar {
 
            min-height: 70px;
            color: white;
            display: flex;
            flex-direction: row;
            justify-content: flex-end;
            align-items: center;
            padding: 0 20px;
            font-size: 30px;
            margin: 0 auto;
            width: 200px;
        }

        #top-bar-logo {
            width: 40px;
            height: 40px;
            margin-right: 5px;
        }

        .container {
            padding-right: 15px;
            padding-left: 15px;
            margin-right: auto;
            margin-left: auto;
        }

        .main-pure-g {
            padding-top: 5vh;
        }

        #links-container {
            display: flex;
            flex-direction: column;
            flex-wrap: wrap;
        }

        #links-container a {
            text-decoration: none;
        }

        .link-item-container {
            background: rgba(205, 210, 209, 0.76);
            text-align: center;
            padding: 10px;
            margin-bottom: 20px;
            cursor: pointer;
            font-size: 18px;
            color: #777777;
            border-radius: 5px;
            box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2);
            font-weight: bold;
        }

        .link-item-container:hover {

            box-shadow: none;
        }

        #content-container {
            padding: 0 20px;
        }

        .card {
            /* Add shadows to create the "card" effect */
            background-color: white;
            box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2);
            transition: 0.3s;
            border-radius: 5px;
            padding: 10px;
        }

        /* On mouse-over, add a deeper shadow */
        .card:hover {
            box-shadow: 0 8px 16px 0 rgba(0, 0, 0, 0.2);
        }

        #panel-container {
            display: flex;
            flex-direction: row;
            margin-top: 5vh;
        }

        #left-panel {
            flex: 0 0 20%;
        }

        #right-panel {
            flex: 1
        }

        .table-td-title {
            font-size: 18px;
            font-weight: bold;
            padding-right: 10px;
            vertical-align: top;
        }

        .table-td-body {
            font-size: 18px;
            display: flex;
            flex-direction: row;
            align-items: center;
        }

        .table-tr {
            padding-bottom: 10px;
        }

        #profile-photo {
            height: 70px;
            width: 70px;
            border-radius: 35px;

            box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2);
            border: solid 1px white;
            margin-right: 20px;
            cursor: pointer;
        }

        .address-container {
            margin-bottom: 20px;
        }

        @media (min-width: 768px) {
            .container {
                width: 750px;
            }
        }

        @media (min-width: 992px) {
            .container {
                width: 970px;
            }

        }

        @media (min-width: 1200px) {
            .container {
                width: 1170px;
            }
        }

        @media (max-width: 900px) {
            #panel-container {
                display: flex;
                flex-direction: column;
            }

            #links-container {
                display: flex;
                flex-direction: row;

            }

            #content-container {
                padding: 0;
            }

            .link-item-container {
                margin-right: 10px;
            }

            #panel-container {
                margin-top: 2vh;
            }
        }

        .m-b-20 {
            margin-bottom: 20px;
        }

    </style>

    <style>
        .activity-item-container {
            display: flex;
            flex-direction: row;
            justify-content: space-between;
        }

        .activity-body {
            font-size: 18px;
            padding: 10px;
            display: flex;
            flex-direction: row;
            align-items: center;
        }

        .activity-body a {
            color: #16A085;
            margin-left: 10px;
            font-weight: bold;
        }

        .activity-datetime {
            opacity: 0.5;
            font-size: 18px;
            display: flex;
            flex-direction: column;
            justify-content: center;
        }

        .activity-tracking-app {
            font-size: 18px;
            opacity: 0.5;
            margin-right: 10px;
        }

        @media (max-width: 900px) {
            .activity-item-container {
                flex-direction: column-reverse;
            }

            .activity-body {
                padding: 10px 0;
                justify-content: space-between;
            }
        }
    </style>
</head>
<body>

<div id="top-bar">
    <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/a/af/Medical_Find_Sources.svg/150px-Medical_Find_Sources.svg.png" id="top-bar-logo"
         alt="well you"/> well you
</div> 

<!-- BEGIN top, page level, container -->

<div class="container">  <div id="panel-container">
<!-- LEFT pannel -->
<div id="left-panel"> <div id="links-container" class="pure-u-1-5">

    <a href="likes.html">
      <div class="link-item-container">Likes</div>
    </a>

    <a href="saves.html">
      <div class="link-item-container">Saves</div>
    </a>

   <a href="actions.html">
      <div class="link-item-container">Actions</div>
    </a>

  <a href="gallery.html">
      <div class="link-item-container">Gallery</div>
    </a>

  <a href="journals.html">
      <div class="link-item-container">Journals</div>
    </a>

  <a href="questionnaires.html">
      <div class="link-item-container">Questionnaires</div>
    </a>

 <a href="readActivities.html">
      <div class="link-item-container">ReadActivities</div>
    </a>

  <a href="reviews.html">
      <div class="link-item-container">reviews</div>
    </a>

  <a href="timeline.html">
      <div class="link-item-container">Timeline</div>
    </a>

 <a href="articles.html">
      <div class="link-item-container">Articles</div>
    </a>

</div> </div> 

<!-- RIGHT pannel -->
<div id="right-panel">  <div id="content-container"> <div class="card"> <table cellpadding="10px">

    <tr class="table-tr">
        <td class="table-td-title">Name</td>
        <td class="table-td-body">{{.firstName}} {{.lastName}}</td>
    </tr>

    <tr class="table-tr">
        <td class="table-td-title">Email</td>
        <td class="table-td-body">{{.email}}</td>
    </tr>

   <tr class="table-tr">
        <td class="table-td-title">Gender</td>
        <td class="table-td-body">{{.gender}}</td>
    </tr>

    <tr class="table-tr">
        <td class="table-td-title">Date of Birth</td>
        <td class="table-td-body">{{.dob}}</td>
    </tr>
            
   <tr class="table-tr">
        <td class="table-td-title">Location</td>
        <td class="table-td-body">{{.location}}</td>
    </tr>

   <tr class="table-tr">
        <td class="table-td-title">Organization</td>
        <td class="table-td-body">{{.university}}</td>
    </tr>

   <tr class="table-tr">
        <td class="table-td-title">Department</td>
        <td class="table-td-body">{{.department}}</td>
    </tr>

   <tr class="table-tr">
        <td class="table-td-title">Wellness Score</td>
        <td class="table-td-body">{{.wellnessScore}}</td>
    </tr>

   <tr class="table-tr">
        <td class="table-td-title">well you Point</td>
        <td class="table-td-body">{{.wellYouPoint}}</td>
    </tr>

   <tr class="table-tr">
        <td class="table-td-title">Signup Date</td>
        <td class="table-td-body">{{.signupDate}}</td>
    </tr>

   <tr class="table-tr">
        <td class="table-td-title">Questionnaires Date</td>
        <td class="table-td-body">{{.quesAnsDate}}</td>
    </tr>

   <tr class="table-tr">
        <td class="table-td-title">Last Health Entry</td>
        <td class="table-td-body">{{.lastHealthEntry}}</td>
    </tr>

   <tr class="table-tr">
        <td class="table-td-title">Last Journal Entry</td>
        <td class="table-td-body">{{.lastJournalEntry}}</td>
    </tr>

   <tr class="table-tr">
        <td class="table-td-title">Last Review Entry</td>
        <td class="table-td-body">{{.lastReviewEntry}}</td>
    </tr>

</div> </div> </div> </div>


</div>  </div>
<!-- END top, page level, container -->
</body>
</html>
`
