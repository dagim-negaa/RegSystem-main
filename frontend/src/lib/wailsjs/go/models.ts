export namespace main {
	
	export class Admin {
	    ID: number;
	    Username: string;
	    Password: string;
	    // Go type: time
	    CreatedAt: any;
	    // Go type: time
	    UpdatedAt: any;
	
	    static createFrom(source: any = {}) {
	        return new Admin(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Username = source["Username"];
	        this.Password = source["Password"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], null);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Student {
	    ID: number;
	    Name: string;
	    Sex: string;
	    Age: number;
	    Birthday: string;
	    Address: string;
	    BirthPlace: string;
	    NameOfChrist: string;
	    MotherName: string;
	    Kebele: string;
	    HouseNo: string;
	    PhoneNo: string;
	    Email: string;
	    Username: string;
	    PriviesSchool: string;
	    EducationLevel: string;
	    WorkPosition: string;
	    ChristFatherName: string;
	    Location: string;
	    EmergencyName: string;
	    EmergencyPhoneNo: string;
	    // Go type: time
	    CreatedAt: any;
	    // Go type: time
	    UpdatedAt: any;
	
	    static createFrom(source: any = {}) {
	        return new Student(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Name = source["Name"];
	        this.Sex = source["Sex"];
	        this.Age = source["Age"];
	        this.Birthday = source["Birthday"];
	        this.Address = source["Address"];
	        this.BirthPlace = source["BirthPlace"];
	        this.NameOfChrist = source["NameOfChrist"];
	        this.MotherName = source["MotherName"];
	        this.Kebele = source["Kebele"];
	        this.HouseNo = source["HouseNo"];
	        this.PhoneNo = source["PhoneNo"];
	        this.Email = source["Email"];
	        this.Username = source["Username"];
	        this.PriviesSchool = source["PriviesSchool"];
	        this.EducationLevel = source["EducationLevel"];
	        this.WorkPosition = source["WorkPosition"];
	        this.ChristFatherName = source["ChristFatherName"];
	        this.Location = source["Location"];
	        this.EmergencyName = source["EmergencyName"];
	        this.EmergencyPhoneNo = source["EmergencyPhoneNo"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], null);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

