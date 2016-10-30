// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/riveryang/spider/cmd"
	//"github.com/PuerkitoBio/goquery"
	//"log"
	//"github.com/riveryang/spider/codec"
	//"fmt"
	"strings"
)
func main() {
	cmd.Execute()
	//baseLink := "http://share.dmhy.org"
	//doc, err := goquery.NewDocument("http://share.dmhy.org/topics/list/page/1")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//codec := codec.DmhyTopicCodec{}
	//topics, err := codec.Decode(doc, baseLink)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//
	//for _, topic := range topics {
	//	fmt.Println(topic)
	//}
}

func execText(text string) string {
	return strings.Replace(strings.Replace(strings.Trim(text, " "), "\t", "", -1), "\n", "", -1)
}
