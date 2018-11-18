const format = {
    pattern: [3, 3, 4],
    open: ['(', ') ', '-']
};
export interface Modelor {
    ID?: number;
    Created?: string | Date;
    Modified?: string | Date;
}
export class Model {
    ID: number;
    Created: string | Date;
    Modified: string | Date;

    constructor(options: Modelor = {}) {
        this.ID = options.ID || null;
        this.Created = Model.isDate(options.Created) ?
                        options.Created : Model.makeDate(options.Created);
        this.Modified = Model.isDate(options.Modified) ?
                        options.Modified : Model.makeDate(options.Modified);
    }

    static isDate(d): d is Date {
        if (d instanceof Date) {
            return true;
        }
        return false;
    }

    static makeDate(strDate: string | null): Date {
        if (strDate == null) {
            return new Date();
        }
        return new Date(strDate);
    }

    static formatPhone(val, depth: number = 0): string {
        if (val.length === 0) {
            return '';
        }
        const n = val.substring(0, format.pattern[depth]);

        return format.open[depth] + n + this.formatPhone(val.substring(format.pattern[depth]), ++depth);
    }
}
