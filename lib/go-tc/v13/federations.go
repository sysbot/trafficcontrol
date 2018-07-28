package v13

import tc "github.com/apache/trafficcontrol/lib/go-tc"

/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

type CDNFederationResponse struct {
	Response []CDNFederation `json:"response"`
}

type CreateCDNFederationResponse struct {
	Response CDNFederation `json:"response"`
	Alerts   []tc.Alert    `json:"alerts"`
}

type UpdateCDNFederationResponse struct {
	Response CDNFederation `json:"response"`
	Alerts   []tc.Alert    `json:"alerts"`
}

type DeleteCDNFederationResponse struct {
	Alerts []tc.Alert `json:"alerts"`
}

type CDNFederation struct {
	ID          *int    `json:"id" db:"id"`
	CName       *string `json:"cname" db:"cname"`
	Ttl         *int    `json:"ttl" db:"ttl"`
	Description *string `json:"description" db:"description"`

	//omitempty only works with primitive types and pointers
	*DeliveryServiceIDs `json:"deliveryService,omitempty"`
}

type DeliveryServiceIDs struct {
	DsId  *int    `json:"id,omitempty" db:"ds_id"`
	XmlId *string `json:"xmlId,omitempty" db:"xml_id"`
}