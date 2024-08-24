"use client"

import React, { useEffect, useState } from 'react';
import {
	Table,
	TableHeader,
	TableBody,
	TableColumn,
	TableRow,
	TableCell
  } from "@nextui-org/table";
// * Import CSS file, you can use CSS module if you want
// ! Change your CSS inside this file
import './page.css'
import Kols from './Utils/Kols';

const build_URL = process.env.BUILD_URL;

const Page = () => {
	// * Use useState to store Kols from API 
	// ! (if you have more optimized way to store data, PLEASE FEELS FREE TO CHANGE)
	var [Kols, setKols] = useState<Kols[]>([]);

	// * Fetch API over here 
	// * Use useEffect to fetch data from API 
	useEffect(() => {
		const fetchKols = async () => {
			try {
				const response = await fetch(`${build_URL}/kols`);
				const data = await response.json();
				setKols(data);
	
			} catch (error) {
				console.log('Error fetching data: ', error);
			}
		}
		
		fetchKols();
	}, []);

	const columns = [
		{ header: 'KolID', field: 'KolID' },
		{ header: 'UserProfileID', field: 'UserProfileID' },
		{ header: 'Language', field: 'Language' },
		{ header: 'Education', field: 'Education' },
		{ header: 'ExpectedSalary', field: 'ExpectedSalary' },
		{ header: 'ExpectedSalaryEnable', field: 'ExpectedSalaryEnable' },
		{ header: 'ChannelSettingTypeID', field: 'ChannelSettingTypeID' },
		{ header: 'IDFrontURL', field: 'IDFrontURL' },
		{ header: 'IDBackURL', field: 'IDBackURL' },
		{ header: 'PortraitURL', field: 'PortraitURL' },
		{ header: 'RewardID', field: 'RewardID' },
		{ header: 'PaymentMethodID', field: 'PaymentMethodID' },
		{ header: 'TestimonialsID', field: 'TestimonialsID' },
		{ header: 'VerificationStatus', field: 'VerificationStatus' },
		{ header: 'Enabled', field: 'Enabled' },
		{ header: 'ActiveDate', field: 'ActiveDate' },
		{ header: 'Active', field: 'Active' },
		{ header: 'CreatedBy', field: 'CreatedBy' },
		{ header: 'CreatedDate', field: 'CreatedDate' },
		{ header: 'ModifiedBy', field: 'ModifiedBy' },
		{ header: 'ModifiedDate', field: 'ModifiedDate' },
		{ header: 'IsRemove', field: 'IsRemove' },
		{ header: 'IsOnBoarding', field: 'IsOnBoarding' },
		{ header: 'Code', field: 'Code' },
		{ header: 'PortraitRightURL', field: 'PortraitRightURL' },
		{ header: 'PortraitLeftURL', field: 'PortraitLeftURL' },
		{ header: 'LivenessStatus', field: 'LivenessStatus' }
	]

	const renderCellValue = (value: any) => {
		if (value instanceof Date) {
			return value.toLocaleTimeString();
		} else if (typeof value === 'boolean') {
			return value ? 'True' : 'False';
		} else {
			return value;
		}
	}

	return (
		<>
			<Table aria-label='KOL Information'>
				<TableHeader>
					{columns.map((column) => (
						<TableColumn key = {column.field}>{column.header}</TableColumn>
					))}
				</TableHeader>
				<TableBody>
					{Kols.map((kol) => (
						<TableRow key = {kol.KolID}>
							{columns.map((column) => (
								<TableCell key = {column.field}>{renderCellValue(kol[column.field as keyof Kols])}</TableCell>
							))}
						</TableRow>
					))}
				</TableBody>
			</Table>
		</>
	)
};

export default Page;